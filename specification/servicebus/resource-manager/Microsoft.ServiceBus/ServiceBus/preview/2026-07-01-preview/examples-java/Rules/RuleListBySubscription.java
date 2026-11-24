
/**
 * Samples for Rules ListBySubscriptions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Rules/RuleListBySubscription.json
     */
    /**
     * Sample code: RulesListBySubscriptions.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void rulesListBySubscriptions(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getRules().listBySubscriptions("ArunMonocle", "sdk-Namespace-1319", "sdk-Topics-2081",
            "sdk-Subscriptions-8691", null, null, com.azure.core.util.Context.NONE);
    }
}
