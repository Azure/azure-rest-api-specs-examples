
import com.azure.resourcemanager.servicebus.fluent.models.SBAuthorizationRuleInner;
import com.azure.resourcemanager.servicebus.models.AccessRights;
import java.util.Arrays;

/**
 * Samples for Namespaces CreateOrUpdateAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/
     * SBNameSpaceAuthorizationRuleCreate.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceAuthorizationRuleCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getNamespaces()
            .createOrUpdateAuthorizationRuleWithResponse("ArunMonocle", "sdk-Namespace-6914", "sdk-AuthRules-1788",
                new SBAuthorizationRuleInner().withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND)),
                com.azure.core.util.Context.NONE);
    }
}
