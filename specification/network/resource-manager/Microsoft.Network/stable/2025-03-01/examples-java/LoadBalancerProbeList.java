
/**
 * Samples for LoadBalancerProbes List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/LoadBalancerProbeList.json
     */
    /**
     * Sample code: LoadBalancerProbeList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerProbeList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancerProbes().list("testrg", "lb",
            com.azure.core.util.Context.NONE);
    }
}
