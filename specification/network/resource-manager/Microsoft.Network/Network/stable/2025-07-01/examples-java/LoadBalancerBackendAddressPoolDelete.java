
/**
 * Samples for LoadBalancerBackendAddressPools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerBackendAddressPoolDelete.json
     */
    /**
     * Sample code: BackendAddressPoolDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void backendAddressPoolDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerBackendAddressPools().delete("testrg", "lb", "backend",
            com.azure.core.util.Context.NONE);
    }
}
