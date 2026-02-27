
/**
 * Samples for IpamPools List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/IpamPools_List.json
     */
    /**
     * Sample code: IpamPools_List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ipamPoolsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getIpamPools().list("rg1", "TestNetworkManager", null, null, null,
            null, null, com.azure.core.util.Context.NONE);
    }
}
