import com.azure.core.util.Context;

/** Samples for Namespaces DeleteAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceAuthorizationRuleDelete.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleDelete.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void nameSpaceAuthorizationRuleDelete(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager
            .namespaces()
            .deleteAuthorizationRuleWithResponse("5ktrial", "nh-sdk-ns", "RootManageSharedAccessKey", Context.NONE);
    }
}
