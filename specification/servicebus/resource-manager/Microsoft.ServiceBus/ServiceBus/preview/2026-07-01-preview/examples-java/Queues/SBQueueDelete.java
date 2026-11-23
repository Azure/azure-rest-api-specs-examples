
/**
 * Samples for Queues Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Queues/SBQueueDelete.json
     */
    /**
     * Sample code: QueueDelete.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void queueDelete(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getQueues().deleteWithResponse("ArunMonocle", "sdk-Namespace-183", "sdk-Queues-8708",
            com.azure.core.util.Context.NONE);
    }
}
