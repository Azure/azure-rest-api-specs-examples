
import com.azure.resourcemanager.servicebus.fluent.models.SBAuthorizationRuleInner;
import com.azure.resourcemanager.servicebus.models.AccessRights;
import java.util.Arrays;

/**
 * Samples for Namespaces CreateOrUpdateAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/SBNameSpaceAuthorizationRuleCreate.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleCreate.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void
        nameSpaceAuthorizationRuleCreate(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNamespaces().createOrUpdateAuthorizationRuleWithResponse("ArunMonocle",
            "sdk-Namespace-6914", "sdk-AuthRules-1788",
            new SBAuthorizationRuleInner().withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND)),
            com.azure.core.util.Context.NONE);
    }
}
