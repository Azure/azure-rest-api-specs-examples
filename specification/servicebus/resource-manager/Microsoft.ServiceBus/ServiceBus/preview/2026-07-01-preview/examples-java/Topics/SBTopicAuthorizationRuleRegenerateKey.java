
import com.azure.resourcemanager.servicebus.models.KeyType;
import com.azure.resourcemanager.servicebus.models.RegenerateAccessKeyParameters;

/**
 * Samples for Topics RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Topics/SBTopicAuthorizationRuleRegenerateKey.json
     */
    /**
     * Sample code: TopicAuthorizationRuleRegenerateKey.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void
        topicAuthorizationRuleRegenerateKey(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getTopics().regenerateKeysWithResponse("Default-ServiceBus-WestUS", "sdk-Namespace8408",
            "sdk-Topics2075", "sdk-Authrules5067", new RegenerateAccessKeyParameters().withKeyType(KeyType.PRIMARY_KEY),
            com.azure.core.util.Context.NONE);
    }
}
