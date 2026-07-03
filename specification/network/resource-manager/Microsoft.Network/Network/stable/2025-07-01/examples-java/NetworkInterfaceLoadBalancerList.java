
/**
 * Samples for NetworkInterfaceLoadBalancers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkInterfaceLoadBalancerList.json
     */
    /**
     * Sample code: NetworkInterfaceLoadBalancerList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkInterfaceLoadBalancerList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaceLoadBalancers().list("testrg", "nic1",
            com.azure.core.util.Context.NONE);
    }
}
