const grpc = require("@grpc/grpc-js");
var protoLoader = require("@grpc/proto-loader");
const PROTO_PATH = "./user.proto";

const options = {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
};

let packageDefinition = protoLoader.loadSync(PROTO_PATH, options);

const UserService = grpc.loadPackageDefinition(packageDefinition).user.UserService;

const client = new UserService(
    "localhost:50051",
    grpc.credentials.createInsecure()
);

client.ReadAll({}, (error, users) => {
    if (!error) console.log(error)
    console.log(users, "userrr");
});