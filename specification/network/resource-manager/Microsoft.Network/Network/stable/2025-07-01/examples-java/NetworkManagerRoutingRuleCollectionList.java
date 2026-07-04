
/**
 * Samples for RoutingRuleCollections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerRoutingRuleCollectionList.json
     */
    /**
     * Sample code: List routing rule collections.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listRoutingRuleCollections(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRoutingRuleCollections().list("rg1", "testNetworkManager", "myTestRoutingConfig",
            null, null, com.azure.core.util.Context.NONE);
    }
}
