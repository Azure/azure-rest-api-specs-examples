
/**
 * Samples for RoutingRuleCollections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerRoutingRuleCollectionGet.json
     */
    /**
     * Sample code: Gets routing rule collection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getsRoutingRuleCollection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRoutingRuleCollections().getWithResponse("rg1", "testNetworkManager",
            "myTestRoutingConfig", "testRuleCollection", com.azure.core.util.Context.NONE);
    }
}
