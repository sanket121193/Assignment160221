module.exports = (database, Sequelize) => {
  return database.define("patientTable", {
    patientid: {
      type: Sequelize.STRING
    },
    patientfname: {
      type: Sequelize.TEXT
    },
    patientlname: {
      type: Sequelize.BOOLEAN
    },
    patientaddress: {
      type: Sequelize.STRING
    }
  });
};
