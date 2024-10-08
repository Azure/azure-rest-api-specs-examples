
import com.azure.resourcemanager.servicebus.fluent.models.SBAuthorizationRuleInner;
import com.azure.resourcemanager.servicebus.models.AccessRights;
import java.util.Arrays;

/**
 * Samples for Topics CreateOrUpdateAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/
     * SBTopicAuthorizationRuleCreate.json
     */
    /**
     * Sample code: TopicAuthorizationRuleCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void topicAuthorizationRuleCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getTopics().createOrUpdateAuthorizationRuleWithResponse(
            "ArunMonocle", "sdk-Namespace-6261", "sdk-Topics-1984", "sdk-AuthRules-4310",
            new SBAuthorizationRuleInner().withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND)),
            com.azure.core.util.Context.NONE);
    }
}
