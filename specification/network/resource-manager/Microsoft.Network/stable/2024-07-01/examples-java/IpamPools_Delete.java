
/**
 * Samples for IpamPools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/IpamPools_Delete.json
     */
    /**
     * Sample code: IpamPools_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ipamPoolsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getIpamPools().delete("rg1", "TestNetworkManager", "TestPool", null,
            com.azure.core.util.Context.NONE);
    }
}
