
/**
 * Samples for Namespaces GetNetworkRuleSet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/VirtualNetworkRule/EHNetworkRuleSetGet.json
     */
    /**
     * Sample code: NameSpaceNetworkRuleSetGet.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void nameSpaceNetworkRuleSetGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().getNetworkRuleSetWithResponse("ResourceGroup", "sdk-Namespace-6019",
            com.azure.core.util.Context.NONE);
    }
}
