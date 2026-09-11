
/**
 * Samples for Namespaces GetNetworkRuleSet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/VirtualNetworkRules/RelayNetworkRuleSetGet.json
     */
    /**
     * Sample code: NameSpaceNetworkRuleSetGet.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void nameSpaceNetworkRuleSetGet(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.namespaces().getNetworkRuleSetWithResponse("ResourceGroup", "example-RelayNamespace-6019",
            com.azure.core.util.Context.NONE);
    }
}
