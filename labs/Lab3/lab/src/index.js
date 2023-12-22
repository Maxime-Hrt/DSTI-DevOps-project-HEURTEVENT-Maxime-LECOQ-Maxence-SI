const express = require('express')
const userRouter = require('./routes/user')
const bodyParser = require('body-parser')
const expressSwagger = require('express-swagger-generator') // Ajoutez cette ligne

const app = express()
const port = process.env.PORT || 3000

const db = require('./dbClient')
db.on("error", (err) => {
  console.error(err)
})

app.use(bodyParser.urlencoded({
  extended: false
}))
app.use(bodyParser.json())

app.get('/', (req, res) => res.send('Hello World!'))

app.use('/user', userRouter)

// Configuration de Swagger - Ajoutez ces lignes
let options = {
  swaggerDefinition: {
    info: {
      description: 'This is a sample server',
      title: 'Swagger',
      version: '1.0.0',
    },
    host: `localhost:${port}`,
    basePath: '/',
    produces: [
      "application/json",
      "application/xml"
    ],
    schemes: ['http', 'https'],
  },
  basedir: __dirname, //app absolute path
  files: ['./routes/**/*.js'] //Path to the API handle folder
};
expressSwagger(app)(options);

const server = app.listen(port, (err) => {
  if (err) throw err
  console.log("Server listening the port " + port)
})

module.exports = server
