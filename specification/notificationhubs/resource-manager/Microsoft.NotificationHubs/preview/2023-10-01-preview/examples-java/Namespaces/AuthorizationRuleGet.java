
/**
 * Samples for Namespaces GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/AuthorizationRuleGet.json
     */
    /**
     * Sample code: Namespaces_GetAuthorizationRule.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        namespacesGetAuthorizationRule(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().getAuthorizationRuleWithResponse("5ktrial", "nh-sdk-ns", "RootManageSharedAccessKey",
            com.azure.core.util.Context.NONE);
    }
}
