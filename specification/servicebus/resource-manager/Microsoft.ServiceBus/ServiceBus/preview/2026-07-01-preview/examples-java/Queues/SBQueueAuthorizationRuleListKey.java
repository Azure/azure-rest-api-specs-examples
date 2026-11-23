
/**
 * Samples for Queues ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Queues/SBQueueAuthorizationRuleListKey.json
     */
    /**
     * Sample code: QueueAuthorizationRuleListKey.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void queueAuthorizationRuleListKey(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getQueues().listKeysWithResponse("ArunMonocle", "sdk-namespace-7982", "sdk-Queues-2317",
            "sdk-AuthRules-5800", com.azure.core.util.Context.NONE);
    }
}
