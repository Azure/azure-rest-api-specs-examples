
import com.azure.resourcemanager.notificationhubs.fluent.models.SharedAccessAuthorizationRuleResourceInner;
import com.azure.resourcemanager.notificationhubs.models.AccessRights;
import java.util.Arrays;

/**
 * Samples for Namespaces CreateOrUpdateAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/AuthorizationRuleCreateOrUpdate.json
     */
    /**
     * Sample code: Namespaces_CreateOrUpdateAuthorizationRule.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void namespacesCreateOrUpdateAuthorizationRule(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces()
            .createOrUpdateAuthorizationRuleWithResponse(
                "5ktrial", "nh-sdk-ns", "sdk-AuthRules-1788", new SharedAccessAuthorizationRuleResourceInner()
                    .withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND)),
                com.azure.core.util.Context.NONE);
    }
}
