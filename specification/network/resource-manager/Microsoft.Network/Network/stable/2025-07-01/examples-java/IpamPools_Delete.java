
/**
 * Samples for IpamPools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/IpamPools_Delete.json
     */
    /**
     * Sample code: IpamPools_Delete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void ipamPoolsDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpamPools().delete("rg1", "TestNetworkManager", "TestPool", null,
            com.azure.core.util.Context.NONE);
    }
}
