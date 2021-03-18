const { Sequelize, DataTypes } = require('sequelize');
const sequelize = new Sequelize('');

const Group = sequelize.define('Group', {
  // Model attributes are defined here
  id: {
      type: DataTypes.INTEGER,
      allowNull: false,
  },
  name: {
    type: DataTypes.STRING,
    allowNull: false
  },

}, {
  // Other model options go here
});

// `sequelize.define` also returns the model