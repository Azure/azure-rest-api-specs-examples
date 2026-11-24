
import com.azure.resourcemanager.servicebus.fluent.models.SBQueueInner;

/**
 * Samples for Queues CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Queues/SBQueueCreate.json
     */
    /**
     * Sample code: QueueCreate.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void queueCreate(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getQueues().createOrUpdateWithResponse("ArunMonocle", "sdk-Namespace-3174",
            "sdk-Queues-5647", new SBQueueInner().withEnablePartitioning(true), com.azure.core.util.Context.NONE);
    }
}
