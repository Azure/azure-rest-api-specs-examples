
/**
 * Samples for RoutingRuleCollections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerRoutingRuleCollectionDelete.json
     */
    /**
     * Sample code: Deletes an routing rule collection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletesAnRoutingRuleCollection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRoutingRuleCollections().delete("rg1", "testNetworkManager",
            "myTestRoutingConfig", "testRuleCollection", null, com.azure.core.util.Context.NONE);
    }
}
