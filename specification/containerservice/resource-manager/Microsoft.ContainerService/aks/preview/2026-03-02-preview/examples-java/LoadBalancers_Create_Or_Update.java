
import com.azure.resourcemanager.containerservice.fluent.models.LoadBalancerInner;

/**
 * Samples for LoadBalancers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/LoadBalancers_Create_Or_Update.json
     */
    /**
     * Sample code: Create or update a Load Balancer.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        createOrUpdateALoadBalancer(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getLoadBalancers().createOrUpdateWithResponse("rg1", "clustername1", "kubernetes",
            new LoadBalancerInner().withPrimaryAgentPoolName("agentpool1").withAllowServicePlacement(true),
            com.azure.core.util.Context.NONE);
    }
}
