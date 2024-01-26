
import com.azure.core.util.Context;

/** Samples for Queues ListByNamespace. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Queues/
     * SBQueueListByNameSpace.json
     */
    /**
     * Sample code: QueueListByNameSpace.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueListByNameSpace(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getQueues().listByNamespace("ArunMonocle",
            "sdk-Namespace-3174", null, null, Context.NONE);
    }
}
