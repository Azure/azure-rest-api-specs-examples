
/**
 * Samples for RoutingRuleCollections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerRoutingRuleCollectionGet.json
     */
    /**
     * Sample code: Gets routing rule collection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsRoutingRuleCollection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRoutingRuleCollections().getWithResponse("rg1",
            "testNetworkManager", "myTestRoutingConfig", "testRuleCollection", com.azure.core.util.Context.NONE);
    }
}
