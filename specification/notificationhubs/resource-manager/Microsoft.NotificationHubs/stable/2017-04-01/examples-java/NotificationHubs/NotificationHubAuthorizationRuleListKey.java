import com.azure.core.util.Context;

/** Samples for NotificationHubs ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubAuthorizationRuleListKey.json
     */
    /**
     * Sample code: NotificationHubAuthorizationRuleListKey.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubAuthorizationRuleListKey(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager
            .notificationHubs()
            .listKeysWithResponse("5ktrial", "nh-sdk-ns", "nh-sdk-hub", "sdk-AuthRules-5800", Context.NONE);
    }
}
