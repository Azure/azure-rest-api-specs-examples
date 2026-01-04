
/**
 * Samples for SecurityUserRuleCollections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerSecurityUserRuleCollectionDelete.json
     */
    /**
     * Sample code: Deletes a Security User Rule collection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletesASecurityUserRuleCollection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityUserRuleCollections().delete("rg1", "testNetworkManager",
            "myTestSecurityConfig", "testRuleCollection", false, com.azure.core.util.Context.NONE);
    }
}
