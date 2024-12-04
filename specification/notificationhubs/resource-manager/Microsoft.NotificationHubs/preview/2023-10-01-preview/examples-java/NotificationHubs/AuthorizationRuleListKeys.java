
/**
 * Samples for NotificationHubs ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NotificationHubs/AuthorizationRuleListKeys.json
     */
    /**
     * Sample code: NotificationHubs_ListKeys.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        notificationHubsListKeys(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().listKeysWithResponse("5ktrial", "nh-sdk-ns", "nh-sdk-hub", "sdk-AuthRules-5800",
            com.azure.core.util.Context.NONE);
    }
}
