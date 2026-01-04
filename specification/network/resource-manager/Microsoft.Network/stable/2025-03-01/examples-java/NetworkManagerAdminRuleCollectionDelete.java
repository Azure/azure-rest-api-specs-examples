
/**
 * Samples for AdminRuleCollections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerAdminRuleCollectionDelete.json
     */
    /**
     * Sample code: Deletes an admin rule collection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletesAnAdminRuleCollection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAdminRuleCollections().delete("rg1", "testNetworkManager",
            "myTestSecurityConfig", "testRuleCollection", false, com.azure.core.util.Context.NONE);
    }
}
