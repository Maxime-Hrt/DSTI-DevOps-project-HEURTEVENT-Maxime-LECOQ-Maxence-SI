const app = require('../src/index')
const chai = require('chai')
const chaiHttp = require('chai-http')
const db = require('../src/dbClient')

chai.use(chaiHttp)

describe('User REST API', () => {

    beforeEach(() => {
        // Clean DB before each test
        db.flushdb()
    })

    after(() => {
        app.close()
        db.quit()
    })

    describe('POST /user', () => {
        it('create a new user', (done) => {
            const user = {
                username: 'sergkudinov',
                firstname: 'Sergei',
                lastname: 'Kudinov'
            }
            chai.request(app)
                .post('/user')
                .send(user)
                .then((res) => {
                    chai.expect(res).to.have.status(201)
                    chai.expect(res.body.status).to.equal('success')
                    chai.expect(res).to.be.json
                    done()
                })
                .catch((err) => {
                    throw err
                })
        })

        it('pass wrong parameters', (done) => {
            const user = {
                firstname: 'Sergei',
                lastname: 'Kudinov'
            }
            chai.request(app)
                .post('/user')
                .send(user)
                .then((res) => {
                    chai.expect(res).to.have.status(400)
                    chai.expect(res.body.status).to.equal('error')
                    chai.expect(res).to.be.json
                    done()
                })
                .catch((err) => {
                    throw err
                })
        })
    })

    describe('GET /user/:username', () => {
        // TODO Create test for the get method
        it('get a user by username', async () => {
            try {
                const user = {
                    username: 'sergkudinov',
                    firstname: 'Sergei',
                    lastname: 'Kudinov'
                }
                // First create a user
                await chai.request(app)
                    .post('/user')
                    .send(user)

                // Then get the user by username
                const res = await chai.request(app)
                    .get(`/user/${user.username}`)

                chai.expect(res).to.have.status(200)
                chai.expect(res.body.data.firstname).to.equal(user.firstname) // Modifier ici
                chai.expect(res.body.data.lastname).to.equal(user.lastname) // et ici
                chai.expect(res).to.be.json
            } catch (err) {
                throw err
            }
        })

        it('cannot get a user when it does not exist', async () => {
            try {
                const res = await chai.request(app)
                    .get('/user/nonexistentuser')
                chai.expect(res).to.have.status(404)
            } catch (err) {
                throw err
            }
        })
    })

})
