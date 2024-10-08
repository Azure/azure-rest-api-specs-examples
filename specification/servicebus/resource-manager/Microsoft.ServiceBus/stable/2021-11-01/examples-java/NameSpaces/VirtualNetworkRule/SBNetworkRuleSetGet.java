
/**
 * Samples for Namespaces GetNetworkRuleSet.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/
     * VirtualNetworkRule/SBNetworkRuleSetGet.json
     */
    /**
     * Sample code: NameSpaceNetworkRuleSetGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceNetworkRuleSetGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getNamespaces()
            .getNetworkRuleSetWithResponse("ResourceGroup", "sdk-Namespace-6019", com.azure.core.util.Context.NONE);
    }
}
