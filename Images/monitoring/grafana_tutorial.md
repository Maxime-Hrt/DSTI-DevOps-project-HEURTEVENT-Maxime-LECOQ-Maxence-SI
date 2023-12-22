# Grafana Tutorial

## Login to Grafana
First of all connect to the grafana address: `http://<MINIKUBE_IP>:32000/`

![img.png](tutografana1.png)

Then when you are on the login page, use the following credentials:
- username: `admin`
- password: `Lecoq-Heurtevent-Password`

## Add a data source
To add a data source, click on Add your first data source:

![img_1.png](tutografana2.png)

Then select Prometheus:

![img_2.png](tutografana3.png)

Then change the URL to `http://prometheus-service.devops-project.svc.cluster.local:9090` and click on Save & Test:

![img_3.png](tutografana4.png)

Then go to the home page

## Then Create a dashboard

Click on create your first dashboard:

![img_1.png](tutografana2.png)

Then click on add a visualization:

![img_4.png](tutografana5.png)

Then select the Prometheus data source you created earlier:

![img_5.png](tutografana6.png)

Then select the query you want to display for example app_health and click on run queries:

![img_6.png](tutografana7.png)

Then click save dashboard:

![img_7.png](tutografana8.png)

Then give a name to your dashboard and click on save and congratulations you have created your first dashboard




