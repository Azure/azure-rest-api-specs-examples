
/**
 * Samples for LoadBalancers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/LoadBalancers_Delete.json
     */
    /**
     * Sample code: LoadBalancers_Delete.
     * 
     * @param manager Entry point to ContainerOrchestratorRuntimeManager.
     */
    public static void loadBalancersDelete(
        com.azure.resourcemanager.containerorchestratorruntime.ContainerOrchestratorRuntimeManager manager) {
        manager.loadBalancers().deleteByResourceGroupWithResponse(
            "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1",
            "testlb", com.azure.core.util.Context.NONE);
    }
}
