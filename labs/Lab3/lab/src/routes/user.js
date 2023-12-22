// router/user.js
const express = require('express')
const userController = require('../controllers/user')

const userRouter = express.Router()

/**
 * @typedef User
 * @property {string} username.required
 * @property {string} firstname.required
 * @property {string} lastname.required
 */

/**
 * This function creates a new user.
 * @route POST /user
 * @group User - Operations about user
 * @param {User.model} user.body.required - new user info
 * @returns {object} 201 - User created successfully
 * @returns {Error}  400 - Error message
 */
userRouter.post('/', (req, resp) => {
  userController.create(req.body, (err, res) => {
    let respObj;
    if (err) {
      respObj = {
        status: "error",
        msg: err.message
      }
      return resp.status(400).json(respObj)
    }
    respObj = {
      status: "success",
      msg: res
    }
    resp.status(201).json(respObj)
  })
})

/**
 * This function gets a user by username.
 * @route GET /user/{username}
 * @group User - Operations about user
 * @param {string} username.path.required - username of the user to fetch
 * @returns {object} 200 - User data
 * @returns {Error}  404 - User not found
 */
userRouter.get('/:username', (req, resp) => {
  const username = req.params.username;
  userController.get(username, (err, res) => {
    let respObj;
    if (err || !res) {
      respObj = {
        status: "error",
        msg: err ? err.message : "User not found"
      }
      return resp.status(404).json(respObj);
    }
    respObj = {
      status: "success",
      data: res
    }
    resp.status(200).json(respObj);
  })
})

module.exports = userRouter
