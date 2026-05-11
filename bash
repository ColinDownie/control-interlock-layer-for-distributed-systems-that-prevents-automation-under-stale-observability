kubectl create deployment demo --image=nginx
kubectl label deployment demo app=demo
cd control-interlock-layer-for-distributed-systems-that-prevents-automation-under-stale-observability
kubectl create deployment demo --image=nginx
kubectl annotate deployment demo sync-test="$(date)"
