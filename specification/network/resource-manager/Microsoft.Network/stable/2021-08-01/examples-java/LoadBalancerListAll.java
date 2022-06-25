import com.azure.core.util.Context;

/** Samples for LoadBalancers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/LoadBalancerListAll.json
     */
    /**
     * Sample code: List all load balancers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllLoadBalancers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancers().list(Context.NONE);
    }
}
