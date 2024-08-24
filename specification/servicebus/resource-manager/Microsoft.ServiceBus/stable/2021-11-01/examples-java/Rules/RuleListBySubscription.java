
/**
 * Samples for Rules ListBySubscriptions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Rules/
     * RuleListBySubscription.json
     */
    /**
     * Sample code: RulesListBySubscriptions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void rulesListBySubscriptions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getRules().listBySubscriptions("ArunMonocle",
            "sdk-Namespace-1319", "sdk-Topics-2081", "sdk-Subscriptions-8691", null, null,
            com.azure.core.util.Context.NONE);
    }
}
