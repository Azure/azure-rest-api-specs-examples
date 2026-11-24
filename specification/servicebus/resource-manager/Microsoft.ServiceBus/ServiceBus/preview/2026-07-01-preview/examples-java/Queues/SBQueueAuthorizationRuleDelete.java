
/**
 * Samples for Queues DeleteAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Queues/SBQueueAuthorizationRuleDelete.json
     */
    /**
     * Sample code: QueueAuthorizationRuleDelete.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void queueAuthorizationRuleDelete(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getQueues().deleteAuthorizationRuleWithResponse("ArunMonocle", "sdk-namespace-7982",
            "sdk-Queues-2317", "sdk-AuthRules-5800", com.azure.core.util.Context.NONE);
    }
}
