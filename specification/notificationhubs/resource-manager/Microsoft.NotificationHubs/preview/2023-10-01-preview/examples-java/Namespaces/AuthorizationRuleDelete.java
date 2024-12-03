
/**
 * Samples for Namespaces DeleteAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/AuthorizationRuleDelete.json
     */
    /**
     * Sample code: Namespaces_DeleteAuthorizationRule.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        namespacesDeleteAuthorizationRule(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().deleteAuthorizationRuleWithResponse("5ktrial", "nh-sdk-ns", "RootManageSharedAccessKey",
            com.azure.core.util.Context.NONE);
    }
}
