
import com.azure.resourcemanager.servicebus.models.KeyType;
import com.azure.resourcemanager.servicebus.models.RegenerateAccessKeyParameters;

/**
 * Samples for Queues RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Queues/
     * SBQueueAuthorizationRuleRegenerateKey.json
     */
    /**
     * Sample code: QueueAuthorizationRuleRegenerateKey.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueAuthorizationRuleRegenerateKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getQueues().regenerateKeysWithResponse("ArunMonocle",
            "sdk-namespace-7982", "sdk-Queues-2317", "sdk-AuthRules-5800",
            new RegenerateAccessKeyParameters().withKeyType(KeyType.PRIMARY_KEY), com.azure.core.util.Context.NONE);
    }
}
