
import com.azure.resourcemanager.servicebus.models.KeyType;
import com.azure.resourcemanager.servicebus.models.RegenerateAccessKeyParameters;

/**
 * Samples for Queues RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Queues/SBQueueAuthorizationRuleRegenerateKey.json
     */
    /**
     * Sample code: QueueAuthorizationRuleRegenerateKey.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void
        queueAuthorizationRuleRegenerateKey(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getQueues().regenerateKeysWithResponse("ArunMonocle", "sdk-namespace-7982",
            "sdk-Queues-2317", "sdk-AuthRules-5800",
            new RegenerateAccessKeyParameters().withKeyType(KeyType.PRIMARY_KEY), com.azure.core.util.Context.NONE);
    }
}
