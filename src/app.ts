import * as express from 'express';
import * as bodyParser from 'body-parser';

import index from './routes/index';

const app: express.Express = express();

// Middleware
app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());

// Routes
app.use('/', index);

// catch 404 and forward to error handler
app.use((req: express.Request, res: express.Response, next: express.NextFunction) => {
  let err: any = new Error('Not Found');
  err['status'] = 404;
  next(err);
});

// catch 404 and forward to error handler
app.use((req: express.Request, res: express.Response, next: express.NextFunction) => {
  let err: any = new Error('Not Found');
  err['status'] = 404;
  next(err);
});

// error handlers

// development error handler
// will print stacktrace
if (app.get('env') === 'development') {
  app.use((error: any, req: express.Request, res: express.Response, next: express.NextFunction) => {
    res.status(error['status'] || 500);
    res.send({
        'error': {
            'message': error.message,
            'error': error
        }
    });
  });
}

// production error handler
// no stacktraces leaked to user
app.use((error: any, req: express.Request, res: express.Response, next: express.NextFunction) => {
  res.status(error['status'] || 500);
  res.send();

  return null;
});


export default app;