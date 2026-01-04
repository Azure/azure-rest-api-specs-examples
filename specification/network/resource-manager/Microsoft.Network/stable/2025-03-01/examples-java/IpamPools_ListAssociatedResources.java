
/**
 * Samples for IpamPools ListAssociatedResources.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * IpamPools_ListAssociatedResources.json
     */
    /**
     * Sample code: IpamPools_ListAssociatedResources.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ipamPoolsListAssociatedResources(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getIpamPools().listAssociatedResources("rg1", "TestNetworkManager",
            "TestPool", com.azure.core.util.Context.NONE);
    }
}
