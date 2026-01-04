
/**
 * Samples for SecurityUserRuleCollections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerSecurityUserRuleCollectionGet.json
     */
    /**
     * Sample code: Gets security user rule collection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsSecurityUserRuleCollection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityUserRuleCollections().getWithResponse("rg1",
            "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", com.azure.core.util.Context.NONE);
    }
}
