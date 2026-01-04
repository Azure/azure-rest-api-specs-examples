
/**
 * Samples for LoadBalancerProbes Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/LoadBalancerProbeGet.json
     */
    /**
     * Sample code: LoadBalancerProbeGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerProbeGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancerProbes().getWithResponse("testrg", "lb", "probe1",
            com.azure.core.util.Context.NONE);
    }
}
