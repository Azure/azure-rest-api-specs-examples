
/**
 * Samples for Namespaces ListNetworkRuleSets.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/VirtualNetworkRule/SBNetworkRuleSetList.json
     */
    /**
     * Sample code: NameSpaceNetworkRuleSetList.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void nameSpaceNetworkRuleSetList(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNamespaces().listNetworkRuleSets("ResourceGroup", "sdk-Namespace-6019",
            com.azure.core.util.Context.NONE);
    }
}
