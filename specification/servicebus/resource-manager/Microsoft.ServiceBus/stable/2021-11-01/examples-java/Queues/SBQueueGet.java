
/**
 * Samples for Queues Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Queues/SBQueueGet.json
     */
    /**
     * Sample code: QueueGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getQueues().getWithResponse("ArunMonocle",
            "sdk-Namespace-3174", "sdk-Queues-5647", com.azure.core.util.Context.NONE);
    }
}
