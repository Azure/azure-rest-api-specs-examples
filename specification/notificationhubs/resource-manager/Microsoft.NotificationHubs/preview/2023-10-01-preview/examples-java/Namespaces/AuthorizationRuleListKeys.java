
/**
 * Samples for Namespaces ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/AuthorizationRuleListKeys.json
     */
    /**
     * Sample code: Namespaces_ListKeys.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void namespacesListKeys(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().listKeysWithResponse("5ktrial", "nh-sdk-ns", "RootManageSharedAccessKey",
            com.azure.core.util.Context.NONE);
    }
}
