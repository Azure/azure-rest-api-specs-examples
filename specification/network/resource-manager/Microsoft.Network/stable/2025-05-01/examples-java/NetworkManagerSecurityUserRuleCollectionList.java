
/**
 * Samples for SecurityUserRuleCollections List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkManagerSecurityUserRuleCollectionList.json
     */
    /**
     * Sample code: List rule collections in a security configuration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listRuleCollectionsInASecurityConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityUserRuleCollections().list("rg1", "testNetworkManager",
            "myTestSecurityConfig", null, null, com.azure.core.util.Context.NONE);
    }
}
