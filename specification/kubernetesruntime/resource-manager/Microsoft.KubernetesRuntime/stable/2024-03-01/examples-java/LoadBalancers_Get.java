
/**
 * Samples for LoadBalancers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/LoadBalancers_Get.json
     */
    /**
     * Sample code: LoadBalancers_Get.
     * 
     * @param manager Entry point to ContainerOrchestratorRuntimeManager.
     */
    public static void loadBalancersGet(
        com.azure.resourcemanager.containerorchestratorruntime.ContainerOrchestratorRuntimeManager manager) {
        manager.loadBalancers().getWithResponse(
            "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1",
            "testlb", com.azure.core.util.Context.NONE);
    }
}
