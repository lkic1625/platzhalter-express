const { Sequelize, DataTypes } = require('sequelize');
const sequelize = new Sequelize('');

const User = sequelize.define('User', {
  // Model attributes are defined here
  id: {
      type: DataTypes.INTEGER,
      allowNull: false,
  },
  name: {
    type: DataTypes.STRING,
    allowNull: false
  },
  email: {
    type: DataTypes.STRING
    // allowNull defaults to true
  }

}, {
  // Other model options go here
});

// `sequelize.define` also returns the model