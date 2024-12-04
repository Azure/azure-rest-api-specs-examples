
import com.azure.resourcemanager.notificationhubs.models.PolicyKeyResource;
import com.azure.resourcemanager.notificationhubs.models.PolicyKeyType;

/**
 * Samples for NotificationHubs RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NotificationHubs/AuthorizationRuleRegenerateKey.json
     */
    /**
     * Sample code: NotificationHubs_RegenerateKeys.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        notificationHubsRegenerateKeys(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().regenerateKeysWithResponse("5ktrial", "nh-sdk-ns", "nh-sdk-hub",
            "DefaultListenSharedAccessSignature", new PolicyKeyResource().withPolicyKey(PolicyKeyType.PRIMARY_KEY),
            com.azure.core.util.Context.NONE);
    }
}
