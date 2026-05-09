signal, err := r.fetchSRM(sm.Spec.SRMEndpoint)
if err != nil {
	return ctrl.Result{RequeueAfter: 2 * time.Second}, nil
}

G := signal.G

state := "GREEN"
paused := false

if G >= sm.Spec.Threshold {
	state = "RED"
	paused = true
} else if G >= sm.Spec.Threshold*0.5 {
	state = "YELLOW"
}

// list deployments
var deployments appsv1.DeploymentList
_ = r.List(ctx, &deployments, client.InNamespace(sm.Spec.Namespace))

for _, d := range deployments.Items {

	match := true
	for k, v := range sm.Spec.Selector {
		if d.Labels[k] != v {
			match = false
		}
	}

	if !match {
		continue
	}

	if d.Spec.Paused != paused {
		d.Spec.Paused = paused
		_ = r.Update(ctx, &d)
	}
}

sm.Status.G = G
sm.Status.State = state
sm.Status.LastEvaluated = time.Now().String()
_ = r.Status().Update(ctx, &sm)

return ctrl.Result{RequeueAfter: 2 * time.Second}, nil
