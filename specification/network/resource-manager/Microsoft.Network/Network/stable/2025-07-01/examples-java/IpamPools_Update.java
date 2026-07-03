
/**
 * Samples for IpamPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/IpamPools_Update.json
     */
    /**
     * Sample code: IpamPools_Update.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void ipamPoolsUpdate(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpamPools().updateWithResponse("rg1", "TestNetworkManager", "TestPool", null, null,
            com.azure.core.util.Context.NONE);
    }
}
