
/**
 * Samples for RoutingRuleCollections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerRoutingRuleCollectionDelete.json
     */
    /**
     * Sample code: Deletes an routing rule collection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deletesAnRoutingRuleCollection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRoutingRuleCollections().delete("rg1", "testNetworkManager", "myTestRoutingConfig",
            "testRuleCollection", null, com.azure.core.util.Context.NONE);
    }
}
