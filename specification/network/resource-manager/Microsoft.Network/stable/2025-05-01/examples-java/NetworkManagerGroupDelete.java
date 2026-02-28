
/**
 * Samples for NetworkGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerGroupDelete.
     * json
     */
    /**
     * Sample code: NetworkGroupsDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkGroupsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkGroups().delete("rg1", "testNetworkManager",
            "testNetworkGroup", false, com.azure.core.util.Context.NONE);
    }
}
