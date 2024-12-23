
/**
 * Samples for RoutingRuleCollections List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/
     * NetworkManagerRoutingRuleCollectionList.json
     */
    /**
     * Sample code: List routing rule collections.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRoutingRuleCollections(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRoutingRuleCollections().list("rg1", "testNetworkManager",
            "myTestRoutingConfig", null, null, com.azure.core.util.Context.NONE);
    }
}
