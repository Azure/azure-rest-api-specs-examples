import com.azure.resourcemanager.notificationhubs.fluent.models.SharedAccessAuthorizationRuleProperties;
import com.azure.resourcemanager.notificationhubs.models.AccessRights;
import java.util.Arrays;

/** Samples for Namespaces CreateOrUpdateAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceAuthorizationRuleCreate.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleCreate.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void nameSpaceAuthorizationRuleCreate(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager
            .namespaces()
            .defineAuthorizationRule("sdk-AuthRules-1788")
            .withExistingNamespace("5ktrial", "nh-sdk-ns")
            .withProperties(
                new SharedAccessAuthorizationRuleProperties()
                    .withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND)))
            .create();
    }
}
