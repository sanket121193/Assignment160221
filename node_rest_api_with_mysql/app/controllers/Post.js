const db = require('../models')
const postObj = db.posts
const Op = db.Sequelize.Op

// Create and save new Post
exports.create = (request, result) => {
  if (!request.body.patientId) {
    result.status(400).send({
      message: "Content cannot be empty"
    });
  }

  // Create a Post object
  const post = {
    patientId: request.body.patientId,
    patientfname: request.body.patientfname,
    patientlname: request.body.published ? request.body.patientlname,
    patientAddress: request.body.publisher ? request.body.patientAddress
  };

  // Save Post object to db
  postObj.create(post).then(data => {
    result.send(data);
  }).catch(err => {
    result.status(500).send({
      message: err.message || "Some error occurred while saving."
    });
  });
};


// Get all published post by PatientID
exports.getPostBypatientId = (request, result) => {
  const patientId = request.query.patientId;
  postObj.findAll({
    where: {
      patientfname: { \[Op.like\]: `%${patientId}%` },
      patientlname: true
    }
  }).then(data => {
    result.send(data);
  }).catch(err => {
    result.status(500).send({
      message: err.message || "Something going wrong. Unable to retrieve data!"
    });
  });
};



// Get Post object by ID
exports.getPostByID = (request, result) => {
  const paramID = request.params.id;
  console.log(paramID);
  console.log(paramID);
  postObj.findAll({
    where: { id: paramID }
  }).then(data => {
    result.send(data);
  }).catch(err => {
    result.status(500).send({
      message: err.message || `Some error occurred while retrieving data with id : ${paramID}`
    });
  });
};


// Update a Post object by the id
exports.updatePostByID = (request, result) => {
  const id = request.params.id;
  postObj.update(request.body, {
    where: { id: id }
  }).then(num => {
    if (num === 1) {
      result.send({
        message: "Post object successfully updated."
      });
    } else {
      result.send({
        message: `Cannot update Post object with id=${id}!`
      });
    }
  }).catch(err => {
    result.status(500).send({
      message: err.message || `Error while updating Post object with id=${id}!`
    });
  });
};



// Delete Post object by ID
exports.deletePostByID = (request, result) => {
  const id = request.params.id;
  postObj.destroy({
    where: { id: id }
  }).then(num => {
    if (num === 1) {
      result.send({
        message: "Post object successfully deleted."
      });
    } else {
      result.send({
        message: `Cannot delete Post object with id=${id}!`
      });
    }
  }).catch(err => {
    result.status(500).send({
      message: err.message || `Cannot delete Post object with id=${id}!`
    });
  });
};
