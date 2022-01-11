package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"

	blog "github.com/example/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	port = 8081
)

var serverCtx context.Context
var serverCtxCancel context.CancelFunc

var blogCollection *mongo.Collection

type BlogItem struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

type Server struct {
}

func (s *Server) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.CreateBlogResponse, error) {
	log.Println("Creating a Blog")

	reqItem := req.GetBlog()

	// Insert a new Blog Item - ignore an ID - will added automatically
	result, err := blogCollection.InsertOne(ctx, BlogItem{
		AuthorId: reqItem.GetAuthorId(),
		Content:  reqItem.GetContent(),
		Title:    reqItem.GetTitle(),
	})

	if err != nil {
		errMsg := fmt.Sprintf("Internal error : %v", err)
		return nil, status.Errorf(codes.Internal, errMsg)
	}

	// Assert type interface of InsertedId into Object Id
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return &blog.CreateBlogResponse{
			Result: &blog.Blog{
				Id:       oid.Hex(),
				AuthorId: reqItem.AuthorId,
				Title:    reqItem.Title,
				Content:  reqItem.Content,
			},
		}, nil
	} else {
		errMsg := fmt.Sprintf("Blog id is invalid. It's not an object id")
		return nil, status.Errorf(codes.Internal, errMsg)
	}
}

func (s *Server) ReadBlog(ctx context.Context, req *blog.ReadBlogRequest) (*blog.ReadBlogResponse, error) {
	log.Println("Reading a Blog")

	blogId := req.GetBlogId()

	// Check if ID is an object id
	oid, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Blog id is invalid. It's not an object id")
	}

	// Filter parameter must be a document containing *query operators* and can be used to select the document
	filter := bson.M{"_id": oid} // similar to - where _id eq %oid%

	result := blogCollection.FindOne(ctx, filter)
	if result.Err() != nil {
		errMsg := fmt.Sprintf("Blog with id = %v is not found", blogId)
		return nil, status.Errorf(codes.NotFound, errMsg)
	}

	// Remember : Unmarshal used pointer to populate the data
	blogItem := &BlogItem{}
	if err = result.Decode(blogItem); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to unmarshall Blog Item")
	}

	return &blog.ReadBlogResponse{
		Result: fromItemToBlog(blogItem),
	}, nil
}

func (s *Server) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.UpdateBlogResponse, error) {
	log.Println("Update a Blog")

	reqItem := req.GetBlog()

	oid, err := primitive.ObjectIDFromHex(reqItem.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Blog id is invalid. It's not an object id")
	}

	filter := bson.M{"_id": oid}

	// result := blogCollection.FindOne(ctx, filter)
	// if result.Err() != nil {
	// 	errMsg := fmt.Sprintf("Blog with id = %v is not found", reqItem.GetId())
	// 	return nil, status.Errorf(codes.NotFound, errMsg)
	// }

	// blogItem := &BlogItem{}
	// if err = result.Decode(blogItem); err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Failed to unmarshall Blog Item")
	// }

	// // Update Blog Item From Request Item

	// blogItem.AuthorId = reqItem.GetAuthorId()
	// blogItem.Title = reqItem.GetTitle()
	// blogItem.Content = reqItem.GetContent()

	// _, err = blogCollection.ReplaceOne(ctx, filter, blogItem)
	// if err != nil {
	// 	errMsg := fmt.Sprintf("Internal error : %v", err)
	// 	return nil, status.Errorf(codes.Internal, errMsg)
	// }

	/* Alternative */

	result := blogCollection.FindOneAndReplace(ctx, filter, &BlogItem{
		AuthorId: reqItem.GetAuthorId(),
		Title:    reqItem.GetTitle(),
		Content:  reqItem.GetContent(),
	})
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			errMsg := fmt.Sprintf("Blog with id = %v is not found", reqItem.GetId())
			return nil, status.Errorf(codes.NotFound, errMsg)
		} else {
			errMsg := fmt.Sprintf("Internal error : %v", result.Err())
			return nil, status.Errorf(codes.Internal, errMsg)
		}
	}

	blogItem := &BlogItem{}
	if err = result.Decode(blogItem); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to unmarshall Blog Item")
	}

	// Update Blog Item From Request Item

	blogItem.AuthorId = reqItem.GetAuthorId()
	blogItem.Title = reqItem.GetTitle()
	blogItem.Content = reqItem.GetContent()

	return &blog.UpdateBlogResponse{
		Result: fromItemToBlog(blogItem),
	}, nil
}

func (s *Server) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) (*blog.DeleteBlogResponse, error) {
	log.Println("Delete a Blog")

	blogId := req.GetBlogId()

	oid, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Blog id is invalid. It's not an object id")
	}

	filter := bson.M{"_id": oid}
	result, err := blogCollection.DeleteOne(ctx, filter)
	if err != nil {
		errMsg := fmt.Sprintf("Internal error : %v", err)
		return nil, status.Errorf(codes.Internal, errMsg)
	}

	if result.DeletedCount > 0 {
		return &blog.DeleteBlogResponse{
			Result: fmt.Sprintf("Succeed to delete Blog with id = %v", blogId),
		}, nil
	}

	errMsg := fmt.Sprintf("Failed to delete Blog with id = %v", blogId)
	return nil, status.Errorf(codes.NotFound, errMsg)
}

func (s *Server) ListBlog(req *blog.ListBlogRequest, stream blog.BlogService_ListBlogServer) error {
	log.Println("List of Blogs")

	ctx := serverCtx

	// Blank filter
	// filter := bson.M{}

	// bson.D vs bson.M
	// bson.D order does 	matters
	// bson.M order doesn't matters
	// Ex :
	// Using bson.D in sort which means the order of the sort is important

	// Sorting base on Request, recommend using bson.D
	findOpts := options.Find()
	if len(req.GetSortBy()) > 0 {
		sort := 1
		if req.GetSort() == "asc" {
			sort = 1
		} else if req.GetSort() == "desc" {
			sort = -1
		}

		sortArgs := strings.Split(req.GetSortBy(), ",")

		sortFilter := bson.D{}
		for _, sortArg := range sortArgs {
			sortFilter = append(sortFilter, bson.E{
				Key:   sortArg,
				Value: sort,
			})
		}

		findOpts.SetSort(sortFilter)
	}

	// Filter base on Request, can using either bson.D or bson.M
	filter := bson.D{}
	if len(req.GetTitle()) > 0 {
		filter = append(filter, bson.E{
			Key: "title",
			Value: bson.D{
				{Key: "$eq", Value: req.GetTitle()},
			},
		})
	}
	if len(req.GetAuthorId()) > 0 {
		filter = append(filter, bson.E{
			Key: "author_id",
			Value: bson.D{
				{Key: "$eq", Value: req.GetAuthorId()},
			},
		})
	}

	cursor, err := blogCollection.Find(ctx, filter, findOpts)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Internal error : %v", err))
	}

	// 1. Reading all documents by once
	// blogs := []BlogItem{}
	// if err := cursor.All(ctx, &blogs); err != nil { // cursor.All - (1) Get all documents, (2) Unmarshall document into Blog Item, and (3) Closed the Cursor
	// 	return status.Errorf(codes.Internal, fmt.Sprintf("Can't retrieve all blogs from cursor : %v", err))
	// }

	// for _, item := range blogs {
	// 	stream.Send(&blog.ListBlogResponse{
	// 		Blog: &blog.Blog{
	// 			Id:       item.Id.Hex(),
	// 			AuthorId: item.AuthorId,
	// 			Title:    item.Title,
	// 			Content:  item.Content,
	// 		},
	// 	})

	// // time.Sleep(3 * time.Second)
	// }

	// 2. Reading all documents by iterate over document
	for cursor.Next(ctx) { // cursor.Next - Move into next document
		item := &BlogItem{}
		if err := cursor.Decode(item); err != nil { // cursor.Decode - Unmarshall current document into Blog Item
			return status.Errorf(codes.Internal, "Failed to unmarshall Blog Item")
		}

		stream.Send(&blog.ListBlogResponse{
			Blog: fromItemToBlog(item),
		})

		// time.Sleep(3 * time.Second)
	}
	cursor.Close(ctx) // Closed the cursor

	return nil
}

func init() {
	serverCtx, serverCtxCancel = context.WithTimeout(context.Background(), time.Duration(15)*time.Minute)
}

func main() {
	// Set log level
	log.SetFlags(log.LstdFlags | log.Lshortfile) // std - standard log (Ldate and Ltime) and shortFile (line number in file)

	// DB Connection
	dbClient := openDB("mongodb+srv://andrewgun95:andregokil@andrewgun.bbygi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	// Defer - Closing the DB
	defer closeDB(dbClient)

	// Setup DB ...
	blogCollection = dbClient.Database("mydb").Collection("blog")

	// Get list of databases
	databaseNames, _ := dbClient.ListDatabaseNames(serverCtx, bson.M{})
	log.Printf("Available databases : %v\n", databaseNames)

	log.Println("Running gRPC server ... at port", port)
	// Create a TCP connection
	lis, err := net.Listen("tcp", fmt.Sprint("127.0.0.1:", port))
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	opts := []grpc.ServerOption{}
	// Create a new gRPC Server
	s := grpc.NewServer(opts...)
	// Defer - Terminating the server
	defer func() {
		log.Println("Terminating gRPC Server ...")
		s.Stop()

		err = lis.Close()
		if err != nil {
			log.Println("TCP listener already closed")
		} else {
			log.Println("Closing TCP listener ...")
		}

		// Cancel the server context
		serverCtxCancel()
	}()
	// Register service into the Server
	blog.RegisterBlogServiceServer(s, &Server{})
	// Register reflection service on gRPC server
	reflection.Register(s)

	// Register service ...

	go func() {
		log.Println("Server ready to serve ...")
		// Bind a TCP connection into gRPC server
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve : %v", err)
		}
	}()

	/* Gracefully Server Termination */
	signals := make(chan os.Signal, 1)

	// Register signal channel to recieved any notification *Unix Signal* interruption - SIGINT (Ctrl + C)
	signal.Notify(signals, os.Interrupt) // signals <- syscall.SIGINT

	// Will wait until recieved a notification from *Unix Signal*
	signal := <-signals
	log.Println("Got unix signal", signal)
}

func openDB(uri string) *mongo.Client {
	log.Println("Open database connection ...")
	errLog := func(err error) { log.Fatalf("Failed to connect database : %v", err) }

	dbClient, dbErr := mongo.NewClient(options.Client().ApplyURI(uri))
	if dbErr != nil {
		errLog(dbErr)
	}
	// Create a connection into database
	dbErr = dbClient.Connect(serverCtx)
	if dbErr != nil {
		errLog(dbErr)
	}
	// Ping to the database !
	if pingErr := dbClient.Ping(serverCtx, readpref.Primary()); pingErr != nil {
		errLog(pingErr)
	}

	return dbClient
}

func closeDB(client *mongo.Client) {
	log.Println("Close database connection ...")

	if err := client.Disconnect(serverCtx); err != nil {
		log.Fatalf("Failed to closed database : %v", err)
	}
}

func fromItemToBlog(blogItem *BlogItem) *blog.Blog {
	return &blog.Blog{
		Id:       blogItem.Id.Hex(),
		AuthorId: blogItem.AuthorId,
		Title:    blogItem.Title,
		Content:  blogItem.Content,
	}
}
