
/**
 * Samples for IpamPools GetPoolUsage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/IpamPools_GetPoolUsage.json
     */
    /**
     * Sample code: IpamPools_GetPoolUsage.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void ipamPoolsGetPoolUsage(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpamPools().getPoolUsageWithResponse("rg1", "TestNetworkManager", "TestPool",
            com.azure.core.util.Context.NONE);
    }
}
