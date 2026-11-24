
/**
 * Samples for Queues ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Queues/SBQueueListByNameSpace.json
     */
    /**
     * Sample code: QueueListByNameSpace.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void queueListByNameSpace(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getQueues().listByNamespace("ArunMonocle", "sdk-Namespace-3174", null, null,
            com.azure.core.util.Context.NONE);
    }
}
