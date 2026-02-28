
/**
 * Samples for IpamPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/IpamPools_Update.json
     */
    /**
     * Sample code: IpamPools_Update.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ipamPoolsUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getIpamPools().updateWithResponse("rg1", "TestNetworkManager",
            "TestPool", null, null, com.azure.core.util.Context.NONE);
    }
}
