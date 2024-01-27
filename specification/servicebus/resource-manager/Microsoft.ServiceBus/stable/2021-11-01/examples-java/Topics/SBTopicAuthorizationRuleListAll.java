
import com.azure.core.util.Context;

/** Samples for Topics ListAuthorizationRules. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/
     * SBTopicAuthorizationRuleListAll.json
     */
    /**
     * Sample code: TopicAuthorizationRuleListAll.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void topicAuthorizationRuleListAll(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getTopics().listAuthorizationRules("ArunMonocle",
            "sdk-Namespace-6261", "sdk-Topics-1984", Context.NONE);
    }
}
