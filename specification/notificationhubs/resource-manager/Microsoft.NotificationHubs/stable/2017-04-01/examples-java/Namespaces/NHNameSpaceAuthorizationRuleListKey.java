
/**
 * Samples for Namespaces ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/
     * NHNameSpaceAuthorizationRuleListKey.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleListKey.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        nameSpaceAuthorizationRuleListKey(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().listKeysWithResponse("5ktrial", "nh-sdk-ns", "RootManageSharedAccessKey",
            com.azure.core.util.Context.NONE);
    }
}
