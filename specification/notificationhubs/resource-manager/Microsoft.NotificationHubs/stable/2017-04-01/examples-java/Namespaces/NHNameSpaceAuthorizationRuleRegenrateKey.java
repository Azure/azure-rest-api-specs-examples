import com.azure.core.util.Context;
import com.azure.resourcemanager.notificationhubs.models.PolicykeyResource;

/** Samples for Namespaces RegenerateKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceAuthorizationRuleRegenrateKey.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleRegenerateKey.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void nameSpaceAuthorizationRuleRegenerateKey(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager
            .namespaces()
            .regenerateKeysWithResponse(
                "5ktrial",
                "nh-sdk-ns",
                "RootManageSharedAccessKey",
                new PolicykeyResource().withPolicyKey("PrimaryKey"),
                Context.NONE);
    }
}
