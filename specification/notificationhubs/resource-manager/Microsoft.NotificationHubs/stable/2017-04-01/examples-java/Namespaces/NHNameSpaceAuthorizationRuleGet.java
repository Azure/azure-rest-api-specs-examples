
/**
 * Samples for Namespaces GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/
     * NHNameSpaceAuthorizationRuleGet.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleGet.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        nameSpaceAuthorizationRuleGet(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().getAuthorizationRuleWithResponse("5ktrial", "nh-sdk-ns", "RootManageSharedAccessKey",
            com.azure.core.util.Context.NONE);
    }
}
