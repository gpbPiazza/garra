## 1. Install the Heroku Monorepo Buildpack

Heroku provides a special buildpack specifically for monorepos that allows you to specify which subdirectory to deploy:

```bash
heroku buildpacks:add -a garra-api https://github.com/lstoll/heroku-buildpack-monorepo
```

## 2. Set the APP_BASE Environment Variable

Tell the monorepo buildpack which directory contains your application:

```bash
heroku config:set -a garra-api APP_BASE=garra-api
```

## 3. Add the Go Buildpack

Since your API is written in Go, you'll need the Go buildpack:

```bash
heroku buildpacks:add -a garra-api heroku/go
```

## 4. Update Your Procfile

Make sure your Procfile contains the proper command to start your application:

```
web: ./bin/api
```

## 5. Configure Go Modules

Ensure your go.mod file is at the root of the garra-api directory. If you're using Go modules, make sure your module path is correctly set.

## 6. Configure the Heroku Go Build

Set the GO_APP_NAME environment variable to specify the package name:

```bash
heroku config:set -a garra-api GO_APP_NAME=github.com/gpbPiazza/garra/cmd/api
```

## 7. Set the Required Environment Variables

Make sure to set the same environment variables required by your application:

```bash
heroku config:set -a garra-api PORT=8080
# Add any other required environment variables
```

## 8. Deploy Your App

Once everything is set up, you can deploy your app:

```bash
git push heroku main
```

## 9. Monitor Your Deployment

Check the logs to ensure your application started correctly:

```bash
heroku logs --tail -a garra-api
```

## 10. Common Issues

If your app fails to start, check:
- Procfile location and format
- Environment variable configuration
- Build logs for compilation errors
- Correct module paths and imports