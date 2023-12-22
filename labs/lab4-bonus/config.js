const env = require("dotenv")
env.config()

var config = {}

config.endpoint = process.env.ENDPOINT
config.key = process.env.KEY

config.database = {
  id: 'User'
}

config.container = {
  id: 'ToDoList'
}

config.items = {
  Maxime: {
    id: 'Maxime1',
    username: 'Maxime-hrt',
    password: '123',
    email: 'maxime.heurtevent@edu.ece.fr',
    address: {
      state: 'Ile-de-France',
      city: 'Paris',
      country: 'France'
    }
  },
  Maxence: {
    id: 'Maxence1',
    username: 'Maxence-Lcq',
    password: '123',
    email: 'maxence.lecoq@edu.ece.fr',
    address: {
      state: 'Ile-De-France',
      city: 'Nanterre',
      country: 'France'
    }
  }
}

module.exports = config
