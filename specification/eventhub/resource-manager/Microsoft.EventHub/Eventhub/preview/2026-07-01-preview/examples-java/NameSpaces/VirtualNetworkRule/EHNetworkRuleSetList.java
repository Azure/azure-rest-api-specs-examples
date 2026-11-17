
/**
 * Samples for Namespaces ListNetworkRuleSet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/VirtualNetworkRule/EHNetworkRuleSetList.json
     */
    /**
     * Sample code: NameSpaceNetworkRuleSetList.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void nameSpaceNetworkRuleSetList(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().listNetworkRuleSetWithResponse("ResourceGroup", "sdk-Namespace-6019",
            com.azure.core.util.Context.NONE);
    }
}
