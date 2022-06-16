import com.azure.core.util.Context;

/** Samples for LoadBalancerNetworkInterfaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LoadBalancerNetworkInterfaceListSimple.json
     */
    /**
     * Sample code: LoadBalancerNetworkInterfaceListSimple.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerNetworkInterfaceListSimple(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getLoadBalancerNetworkInterfaces()
            .list("testrg", "lb", Context.NONE);
    }
}
