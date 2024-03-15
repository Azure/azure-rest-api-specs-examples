
import com.azure.resourcemanager.notificationhubs.models.AccessRights;
import java.util.Arrays;

/**
 * Samples for NotificationHubs CreateOrUpdateAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NotificationHubs/AuthorizationRuleCreateOrUpdate.json
     */
    /**
     * Sample code: NotificationHubs_CreateOrUpdateAuthorizationRule.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubsCreateOrUpdateAuthorizationRule(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().defineAuthorizationRule("MyManageSharedAccessKey")
            .withExistingNotificationHub("5ktrial", "nh-sdk-ns", "nh-sdk-hub")
            .withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND)).create();
    }
}
