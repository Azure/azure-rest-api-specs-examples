
/**
 * Samples for IpamPools GetPoolUsage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/IpamPools_GetPoolUsage.json
     */
    /**
     * Sample code: IpamPools_GetPoolUsage.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ipamPoolsGetPoolUsage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getIpamPools().getPoolUsageWithResponse("rg1", "TestNetworkManager",
            "TestPool", com.azure.core.util.Context.NONE);
    }
}
