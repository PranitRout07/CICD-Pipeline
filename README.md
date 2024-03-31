<img width="1280" alt="blackbox exporter dashboard 1" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/8c4b30e8-15fe-47ee-b993-bbbc33d9e16e"># STEPS FOLLOWED FOR THE ENTIRE CICD PIPELINE SETUP

#### 1) Created two github repositories. 
One repository to store the code and another repository to store the manifest files . The manifest repo is a private repository .  
#### 2) Started a Jenkins server at port 8080 and installed some plugins . 
<img width="1280" alt="Install Plugins" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/5b21ef7c-1322-4a63-a3dc-cb4c21c8bbde">

#### 3) Started a sonarqube server at port 9000 . 
<img width="1279" alt="start sonarqube server" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/f80fedf7-0237-4845-8328-9d81d18c5880">

#### 4) Created a token on Sonarqube server ( at adminstration > security > users ). 
<img width="1018" alt="create token on sonarqube" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/6fb85bf4-065a-481e-b691-5bdecc55078f">

#### 5) Configure sonarqube server on Jenkins . 
Here mentioned the sonarqube server URL and also created a secret text using the token ( created in step 4 ) and added that secret text into the sonarqube server . 
<img width="1165" alt="configure sonarqube server on jenkins" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/45c673ed-801d-4efb-abf2-2a1896be7e5a">

#### 6) Setup Sonarqube scanner on Jenkins . 
Here in tools section of manage Jenkins i have configured the sonarqube scanner . Just given the desired version of sonarqube scanner and a name to the scanner . 
<img width="1189" alt="sonarscanner setup on jenkins" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/5240ed95-2994-47b5-b3d5-b770770e6ef2">

#### 7) Setup Quality gate on sonarqube . 
On Sonarqube server , at adminstration ? configuration > webhook , added the <jenkins_url>/sonarqube_webhook and a name . 
<img width="839" alt="quality gate configuration" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/e44ad94a-6eef-4381-92b5-8dbc70d4a8d2">

#### 8) Configure gmail and jenkins integration . 
Here i have first created an app password on gmail . 
<img width="539" alt="created app password on gmail" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/5dc50dd5-87cd-4b98-b458-23299dba693f">

Then using the app password , created a username and password credential on Jenkins . 
<img width="1274" alt="create id for mail-cred" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/27446fff-0e53-4c31-bc98-1da3773b8813">

Then on system of manage jenkins configured the gmail . Here i have mentioned the smtp server and its port as 465 . 
<img width="1280" alt="email configuration on jenkins" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/70ae3384-1741-4605-82ed-721b8995b376">

#### 9) Integrated dockerhub and Jenkins . 
Here created a username and password credential on the Jenkins server . Here i have mentioned my dockerhub account username and the password . Here the username and password can be
accessed by using the ID . 
<img width="1152" alt="create dockerhub credentials" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/db63374e-ff33-4c25-95e3-c021b247b664">

#### 10) Integrated Github and Jenkins . 
Here first created a personal token on github . 
<img width="1280" alt="created a token on github" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/faf2cbab-96a1-408b-9c5d-f4ac1fb17129">

Then created a secret text credential by using the github personal token on Jenkins . 
<img width="1142" alt="github credential on jenkins" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/de8e4031-e9f2-46c5-99c6-7198de6ecb5a">

#### 11) Created a K8s cluster using Kind . 
```
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraPortMappings:
  - containerPort: 32215
    hostPort: 80
  - containerPort: 32216
    hostPort: 8081
- role: worker
```
<img width="458" alt="created a cluster" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/e41f2c82-50af-4b20-a745-ab8adca1fbe6">

#### 12) ArgoCD setup using the private manifest repository . 
Here first created argocd namespace . 
<img width="659" alt="namespace and argocd configured" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/13d896e7-c0b3-4e69-8776-0f71f2356ef2">

Then generated the argocd server password . 
<img width="573" alt="argocd server password" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/6033589f-9add-415f-b62a-67a20df96692">

Then port forward the argocd server access to port 8078 . 
<img width="375" alt="port-forward to 8078" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/07042186-3cd9-4c81-908a-2173b2aa231a">

After connecting the argocd server , i have first updated the argocd password . 
<img width="1280" alt="updated argo password" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/3e8284d5-4336-4b5b-9983-4ff36dd85950">

Then connected the manifest repo file using the github manifest repo url , github username and github personal token . 
<img width="1280" alt="connected to manifest repo" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/d5fc9003-fc28-423c-8498-0e8813fc0f68">

Created an application from the connected github repo . 
<img width="1280" alt="created weather app and also give namespace" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/eafcd7e5-18c3-430c-aed2-52554471fc86">

After this argocd started deploying the manifest files . 
<img width="1280" alt="argo cd started deploying the images" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/fdfa6509-bd92-4c3d-a064-4ec5e148e559">

#### 13) Written the Jenkins pipeline code . 
<img width="1280" alt="created pipeline" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/cc96980f-055c-4d33-8153-8e569c027585">

#### 14) Started prometheus at port 9090, grafana at port 3000 and blackbox exporter at port 9115 . 
Here i have configured the prometheus.yml inside the prometheus container . Here added the some code which will help to scrape the metrics from the blacbox exported and Jenkins . 
<img width="782" alt="configure prometheus yaml" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/367beee8-bc46-4934-9792-75f73d4732a9">

Then restarted my prometheus container . After this checked the targets in prometheus server . 
<img width="1280" alt="check the targets on prometheus" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/97e2373c-cddd-4b99-abc1-3f3691279c3d">

Then integrated prometheus and grafana . 
<img width="907" alt="grafana connected to prometheus" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/12d15601-b0b5-4ce8-971c-42b146857a47">

#### 15) Now started the my Jenkins pipeline . 
It will build the version 1 of my website . 
###### Jenkins Pipeline result: 
<img width="1280" alt="pipeline successful" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/d08ff2db-6ac6-45c3-b8d2-8ccbc0a0ca4e">

###### Sonarqube result : 
<img width="1280" alt="sonarqube" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/1792d3ae-fa7f-4d5f-bd1e-2dcdfdbec871">

###### Argo CD deploys the version 1 of website :
<img width="1280" alt="argo cd automatically deployed" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/afbb9e50-27b0-453e-b2ce-d154e1e9c42f">

###### Received the trivy reports and pipeline status on my gmail . 
<img width="1180" alt="All the successfull stages is received at email" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/44dec9f5-c580-4950-8ef5-da5fb20e1752">
<img width="1176" alt="trivy scan reports received" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/e2cb1b00-7b3b-4551-bd67-22281ad68d35">


#### 16) Updated the code . 
Now the website of version 2 . 
<img width="659" alt="updated to version 2" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/8bfeb53e-95ad-4d84-872d-6f77784b6d4b">

###### Jenkins Pipeline result after website code is updated to version 2 :
<img width="1280" alt="version 2 pipeline is successful" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/7464e891-6276-4c91-a5ab-5846f42e3f91">

###### Argo CD automatically deploys the version of website . 
![Uploading argo cd version 2 access.png…]()
<img width="1280" alt="argo cd deploying version 2" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/002b99fb-46ad-4780-bbf4-d51c51610dd6">

###### Received status of pipeline on gmail:
<img width="1280" alt="successful stages" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/90d29de6-b143-48e7-a18d-b6e05f6defba">
<img width="1280" alt="updated trivy reports" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/10568a19-bc63-4e3b-b6c1-afbe411175c0">

#### 17) Accessed the version 2 of website . 
<img width="1280" alt="version 2 of website" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/039827c5-89a1-4b46-ab89-71d32dd71d73">

###                                                                             ---ALL STEPS COMPLETED----

# Manifest files
Here kubernetes secrets is used to store the API_TOKEN . Application runs at port 80 and it is connected to the backend at port 8081 . 

# Received which stages are failed on gmail while running a demo pipeline .
<img width="956" alt="pipeline failed and their stages" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/f4da3b21-566d-436b-9acb-94db5ae001a4">

# Monitoring results : 
### Blackbox exporter : 
1) 

<img width="1280" alt="blackbox exporter dashboard 1" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/79fc1d7a-df1f-44e3-a573-1c14716c3b1b">


2) 

<img width="1280" alt="uptime" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/2dfb1e17-cf5c-4709-989d-80f5eab77312">

### Jenkins :
<img width="1280" alt="Jenkins dashboard monitoring 1" src="https://github.com/PranitRout07/CICD-Pipeline/assets/102309095/0b0e5b7c-44ff-44ca-9b58-16b47572319c">



