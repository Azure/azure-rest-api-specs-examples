
/**
 * Samples for Rules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Rules/RuleDelete.json
     */
    /**
     * Sample code: RulesDelete.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void rulesDelete(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getRules().deleteWithResponse("ArunMonocle", "sdk-Namespace-1319", "sdk-Topics-2081",
            "sdk-Subscriptions-8691", "sdk-Rules-6571", com.azure.core.util.Context.NONE);
    }
}
