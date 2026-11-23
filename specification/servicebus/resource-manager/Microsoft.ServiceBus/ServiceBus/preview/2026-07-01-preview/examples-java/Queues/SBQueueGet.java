
/**
 * Samples for Queues Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Queues/SBQueueGet.json
     */
    /**
     * Sample code: QueueGet.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void queueGet(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getQueues().getWithResponse("ArunMonocle", "sdk-Namespace-3174", "sdk-Queues-5647",
            com.azure.core.util.Context.NONE);
    }
}
