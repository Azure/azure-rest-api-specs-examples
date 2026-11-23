
/**
 * Samples for Queues ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Queues/SBQueueAuthorizationRuleListAll.json
     */
    /**
     * Sample code: QueueAuthorizationRuleListAll.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void queueAuthorizationRuleListAll(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getQueues().listAuthorizationRules("ArunMonocle", "sdk-Namespace-7982",
            "sdk-Queues-2317", com.azure.core.util.Context.NONE);
    }
}
