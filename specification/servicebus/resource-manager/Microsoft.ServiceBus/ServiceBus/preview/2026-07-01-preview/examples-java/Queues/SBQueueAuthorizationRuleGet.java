
/**
 * Samples for Queues GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Queues/SBQueueAuthorizationRuleGet.json
     */
    /**
     * Sample code: QueueAuthorizationRuleGet.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void queueAuthorizationRuleGet(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getQueues().getAuthorizationRuleWithResponse("ArunMonocle", "sdk-Namespace-7982",
            "sdk-Queues-2317", "sdk-AuthRules-5800", com.azure.core.util.Context.NONE);
    }
}
