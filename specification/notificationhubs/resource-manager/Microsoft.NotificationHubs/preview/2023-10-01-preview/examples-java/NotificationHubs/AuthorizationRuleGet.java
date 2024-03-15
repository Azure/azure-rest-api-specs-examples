
/**
 * Samples for NotificationHubs GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NotificationHubs/AuthorizationRuleGet.json
     */
    /**
     * Sample code: NotificationHubs_GetAuthorizationRule.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubsGetAuthorizationRule(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().getAuthorizationRuleWithResponse("5ktrial", "nh-sdk-ns", "nh-sdk-hub",
            "DefaultListenSharedAccessSignature", com.azure.core.util.Context.NONE);
    }
}
