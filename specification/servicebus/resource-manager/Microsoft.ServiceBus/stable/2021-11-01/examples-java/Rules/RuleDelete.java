
/**
 * Samples for Rules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Rules/RuleDelete.json
     */
    /**
     * Sample code: RulesDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void rulesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getRules().deleteWithResponse("ArunMonocle",
            "sdk-Namespace-1319", "sdk-Topics-2081", "sdk-Subscriptions-8691", "sdk-Rules-6571",
            com.azure.core.util.Context.NONE);
    }
}
