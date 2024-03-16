
/**
 * Samples for NotificationHubs ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NotificationHubs/AuthorizationRuleList.json
     */
    /**
     * Sample code: NotificationHubs_ListAuthorizationRules.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubsListAuthorizationRules(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().listAuthorizationRules("5ktrial", "nh-sdk-ns", "nh-sdk-hub",
            com.azure.core.util.Context.NONE);
    }
}
