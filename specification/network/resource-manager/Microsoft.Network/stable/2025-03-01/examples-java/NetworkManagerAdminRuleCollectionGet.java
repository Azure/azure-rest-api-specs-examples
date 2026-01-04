
/**
 * Samples for AdminRuleCollections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerAdminRuleCollectionGet.json
     */
    /**
     * Sample code: Gets security admin rule collection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsSecurityAdminRuleCollection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAdminRuleCollections().getWithResponse("rg1",
            "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", com.azure.core.util.Context.NONE);
    }
}
