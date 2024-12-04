
/**
 * Samples for Namespaces ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/AuthorizationRuleList.json
     */
    /**
     * Sample code: Namespaces_ListAuthorizationRules.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        namespacesListAuthorizationRules(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().listAuthorizationRules("5ktrial", "nh-sdk-ns", com.azure.core.util.Context.NONE);
    }
}
