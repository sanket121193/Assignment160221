const post = require("../controllers/Post");
const express = require("express");
const router = express.Router();
// Create New Post
router.post("/api/posts/create", post.create);
// // Retrieve all posts
router.get("/api/posts/all", post.getAllPosts);
// Retrieve all patientname posts
router.get("/api/posts/patientfname", post.getAllPatientsfirstname);
// Retrieve all Patient last name post
router.get("/api/posts/patientlname", post.getAllPatientslastName);
// Retrieve all posts by patient address
router.get("/api/posts", post.getPostBypatientAddress);
// Retrieve post by ID
router.get("/api/posts/:patientid", post.getPostBypatientId);
// // Update post by ID
router.put("/api/post/update/:patientid", post.updatePostBypatientId);
// // Delete post by ID
router.delete("/api/post/delete/:patientid", post.deletePostBypatientId);
// Delete all posts
router.delete("/api/posts/deleteAll", post.deleteAllPosts);

module.exports = router;
