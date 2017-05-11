import { Router } from 'express';

const index: Router = Router();

/* GET home page. */
index.get('/', function(req, res, next) {
  res.send('Hello World!')
});

export default index;