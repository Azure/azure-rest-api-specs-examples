
import com.azure.resourcemanager.servicebus.models.KeyType;
import com.azure.resourcemanager.servicebus.models.RegenerateAccessKeyParameters;

/**
 * Samples for Topics RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/
     * SBTopicAuthorizationRuleRegenerateKey.json
     */
    /**
     * Sample code: TopicAuthorizationRuleRegenerateKey.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void topicAuthorizationRuleRegenerateKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getTopics().regenerateKeysWithResponse(
            "Default-ServiceBus-WestUS", "sdk-Namespace8408", "sdk-Topics2075", "sdk-Authrules5067",
            new RegenerateAccessKeyParameters().withKeyType(KeyType.PRIMARY_KEY), com.azure.core.util.Context.NONE);
    }
}
