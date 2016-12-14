what is nginxConfdForKubernets(nck) ？
    when we setup an kubernets cluster,we found that ,all service just supply for the cluster only .but we should supply for external programe or users,how to deal it ?

    Kubenet support three models(Loadblances,clusterFirst,Nodeport) to expose the service to external and we know,the "Loadblances" just for cloud, the "Nodeport" show open the Mother Host for containers(means that we show expose the Mother Host ,I don't like that ,because dangerous).

   For Minimal change principle,I recommand "clusterFirt" model,and then use proxy(recommand nginx) mode to supply external visit.but when container's IP distribution random,So how to stay in step between kubernet and nginx upstream servers list ?

  Here,the nck can help you deal it!

  so,the Prerequisite are 
   1) the proxy host(nginx) must can visit the containers IP.
   2) the external network can visit the proxy host(nginx).
   3) the nginx version must support tcp/udp proxy (you can get it at http://nginx.org)


How to use it ?

1 download the binary program(Linux 7 x86_64) at realses page or clone the source code build by yourself.
2 edit the config file app.conf and Particular attention :

	#the nginx config file template path the example at
	# https://github.com/jetlwx/nginxConfdForKubernets/tree/master/example
	nginxTemplate = "/etc/nginx/nginx.conf.tmpl"

	#the nginx config file path that for nginx binary ,Depending on your nginx program
	nginxConfFile = "/etc/nginx/nginx.conf"

	#the kubernets api server or etcd server(recommand kubernets)
	apiServer="http://172.16.6.160:8080"

	#list is  namespacename/endpoints/serviceName  ,and split use ","
	#eg: default/endpoints/redis1, kube-system/endpoints/kube-dns
	servicelist="default/endpoints/redis1,default/endpoints/tomcattest"

	#nginx check command
	checkCmd = " /usr/sbin/nginx -t"

	#nginx reload command
	reloadCmd = "/usr/sbin/nginx -s reload"

	# how long (seconds) fresh the endpoints  
	freshSeconds = 10

3 the template file explain
  please see the  template file comment.

