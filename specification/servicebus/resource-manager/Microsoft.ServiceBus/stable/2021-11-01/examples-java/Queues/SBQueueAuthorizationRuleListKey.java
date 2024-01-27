
import com.azure.core.util.Context;

/** Samples for Queues ListKeys. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Queues/
     * SBQueueAuthorizationRuleListKey.json
     */
    /**
     * Sample code: QueueAuthorizationRuleListKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueAuthorizationRuleListKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getQueues().listKeysWithResponse("ArunMonocle",
            "sdk-namespace-7982", "sdk-Queues-2317", "sdk-AuthRules-5800", Context.NONE);
    }
}
