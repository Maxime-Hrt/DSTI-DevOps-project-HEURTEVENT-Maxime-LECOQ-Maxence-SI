// controller/user.js
const db = require('../dbClient')

module.exports = {
  create: (user, callback) => {
    // Check parameters
    if(!user.username)
      return callback(new Error("Wrong user parameters"), null)

    // TODO Verify if the user already exists
    db.hgetall(user.username, (err, existingUser) => {
      if (err) return callback(err, null);
      if (existingUser && Object.keys(existingUser).length) {
        return callback(new Error("User already exists"), null);
      }

      // Create User schema
      const userObj = {
        firstname: user.firstname,
        lastname: user.lastname,
      }

      // Save to DB
      db.hmset(user.username, userObj, (err, res) => {
        if (err) return callback(err, null)
        callback(null, res) // Return callback
      })
    });
  },
  get: (username, callback) => {
    // TODO create this method
    db.hgetall(username, (err, result) => {
      if (err) return callback(err, null);
      callback(null, result);
    });
  }
}
