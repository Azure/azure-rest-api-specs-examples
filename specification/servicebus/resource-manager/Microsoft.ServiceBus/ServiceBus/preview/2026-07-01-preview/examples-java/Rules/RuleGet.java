
/**
 * Samples for Rules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Rules/RuleGet.json
     */
    /**
     * Sample code: RulesGet.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void rulesGet(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getRules().getWithResponse("ArunMonocle", "sdk-Namespace-1319", "sdk-Topics-2081",
            "sdk-Subscriptions-8691", "sdk-Rules-6571", com.azure.core.util.Context.NONE);
    }
}
