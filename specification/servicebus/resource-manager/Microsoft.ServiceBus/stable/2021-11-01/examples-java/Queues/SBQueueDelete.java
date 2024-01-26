
import com.azure.core.util.Context;

/** Samples for Queues Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Queues/SBQueueDelete.
     * json
     */
    /**
     * Sample code: QueueDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getQueues().deleteWithResponse("ArunMonocle",
            "sdk-Namespace-183", "sdk-Queues-8708", Context.NONE);
    }
}
