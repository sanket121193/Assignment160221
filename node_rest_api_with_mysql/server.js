const express = require("express");
const bodyParser = require("body-parser");
const cors = require("cors");
const server = express();
const db = require("./models");
const api = require("./routes/index");
const corsSettings = {
  originL: "http://localhost:8081"
};
server.use(cors(corsSettings));

const db = require("./models");

// Parse request of content-type - application/json
server.use(bodyParser.json());

// parse requests of content-type -application/x-www-form-urlencoded
server.use(bodyParser.urlencoded({ extended: true }));

server.use("/", api);


//create a simple route
server.get("/", (\_req, res) => {
   res.json({ message: "Node app started" });
});

// set listening ports for request
const port = process.env.PORT || 8080;

server.listen(port, () => {
  console.log("Server running on port : " + port );
});
db.databaseConf.sync();


