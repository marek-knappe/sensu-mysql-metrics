Sensu Mysql Metrics
===================
Golang sensu/graphite plugin to get all metrics from mysql 'show global status'


Why new plugin
===================
I'm sick of ruby plugins that require a lot of dependiences, so I made my script in goLang to get all variables from mysql show global status to sensu (and then to graphite).

You are high-school "developer" that doesn't know what is golang
=======================================================
Yep, It's my first Go program, so it's not the nicest - but it works

Compile and usage
===================
Compile is simple:

```
git build sensu-mysql-metrics.go
```

Usage is even simpler:
```
./sensu-mysql-metrics -username="monitoring" -password="some_password"
```

For more options just use 
```
./sensu-mysql-metrics --help
```

Example sensu file:
```
{
  "checks": {
    "mysql_graphite_metrics": {
      "type": "metric",
      "command": "sensu-mysql-metrics -hostname=':::mysql.host:::'  -username=':::mysql.user:::' -password=':::mysql.password:::'",
      "interval": 60,
      "subscribers": ["mysql"],
      "handlers": ["relay"]
    }
  }
}
```
and just add the config variables on client:

```
    "mysql": {
      "host": "localhost",
      "port": 3306,
      "user": "monitoring",
      "password": "some_password"
    }

```

Known issue
==================
If the mysql parameters are wrong there is an error:
```
panic: runtime error: invalid memory address or nil pointer dereference
```
Will fix it in my free time

