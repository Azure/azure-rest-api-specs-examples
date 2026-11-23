
import com.azure.resourcemanager.servicebus.fluent.models.SBAuthorizationRuleInner;
import com.azure.resourcemanager.servicebus.models.AccessRights;
import java.util.Arrays;

/**
 * Samples for Topics CreateOrUpdateAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Topics/SBTopicAuthorizationRuleCreate.json
     */
    /**
     * Sample code: TopicAuthorizationRuleCreate.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void topicAuthorizationRuleCreate(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getTopics().createOrUpdateAuthorizationRuleWithResponse("ArunMonocle",
            "sdk-Namespace-6261", "sdk-Topics-1984", "sdk-AuthRules-4310",
            new SBAuthorizationRuleInner().withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND)),
            com.azure.core.util.Context.NONE);
    }
}
