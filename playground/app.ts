import { Backend } from 'kuzzle'

const app = new Backend('playground')

app.start()
  .then(() => {
    app.log.info('Application started')
  })
  .catch(console.error)
