
/**
 * Samples for AdminRuleCollections List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkManagerAdminRuleCollectionList.json
     */
    /**
     * Sample code: List security admin rule collections.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSecurityAdminRuleCollections(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAdminRuleCollections().list("rg1", "testNetworkManager",
            "myTestSecurityConfig", null, null, com.azure.core.util.Context.NONE);
    }
}
