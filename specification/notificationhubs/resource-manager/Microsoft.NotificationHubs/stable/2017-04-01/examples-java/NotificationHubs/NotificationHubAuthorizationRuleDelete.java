import com.azure.core.util.Context;

/** Samples for NotificationHubs DeleteAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubAuthorizationRuleDelete.json
     */
    /**
     * Sample code: NotificationHubAuthorizationRuleDelete.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubAuthorizationRuleDelete(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager
            .notificationHubs()
            .deleteAuthorizationRuleWithResponse(
                "5ktrial", "nh-sdk-ns", "nh-sdk-hub", "DefaultListenSharedAccessSignature", Context.NONE);
    }
}
