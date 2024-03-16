
/**
 * Samples for NotificationHubs DeleteAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NotificationHubs/AuthorizationRuleDelete.json
     */
    /**
     * Sample code: NotificationHubs_DeleteAuthorizationRule.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubsDeleteAuthorizationRule(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().deleteAuthorizationRuleWithResponse("5ktrial", "nh-sdk-ns", "nh-sdk-hub",
            "DefaultListenSharedAccessSignature", com.azure.core.util.Context.NONE);
    }
}
