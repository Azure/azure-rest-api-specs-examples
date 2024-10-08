
import com.azure.resourcemanager.servicebus.fluent.models.SBAuthorizationRuleInner;
import com.azure.resourcemanager.servicebus.models.AccessRights;
import java.util.Arrays;

/**
 * Samples for Queues CreateOrUpdateAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Queues/
     * SBQueueAuthorizationRuleCreate.json
     */
    /**
     * Sample code: QueueAuthorizationRuleCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueAuthorizationRuleCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getQueues().createOrUpdateAuthorizationRuleWithResponse(
            "ArunMonocle", "sdk-Namespace-7982", "sdk-Queues-2317", "sdk-AuthRules-5800",
            new SBAuthorizationRuleInner().withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND)),
            com.azure.core.util.Context.NONE);
    }
}
