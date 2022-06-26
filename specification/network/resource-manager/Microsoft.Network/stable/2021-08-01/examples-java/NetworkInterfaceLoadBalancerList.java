import com.azure.core.util.Context;

/** Samples for NetworkInterfaceLoadBalancers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkInterfaceLoadBalancerList.json
     */
    /**
     * Sample code: NetworkInterfaceLoadBalancerList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkInterfaceLoadBalancerList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkInterfaceLoadBalancers()
            .list("testrg", "nic1", Context.NONE);
    }
}
