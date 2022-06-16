import com.azure.core.util.Context;

/** Samples for NotificationHubs GetAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubAuthorizationRuleGet.json
     */
    /**
     * Sample code: NotificationHubAuthorizationRuleGet.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubAuthorizationRuleGet(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager
            .notificationHubs()
            .getAuthorizationRuleWithResponse(
                "5ktrial", "nh-sdk-ns", "nh-sdk-hub", "DefaultListenSharedAccessSignature", Context.NONE);
    }
}
