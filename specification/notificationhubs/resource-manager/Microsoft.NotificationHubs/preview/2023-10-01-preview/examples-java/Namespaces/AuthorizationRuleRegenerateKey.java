
import com.azure.resourcemanager.notificationhubs.models.PolicyKeyResource;
import com.azure.resourcemanager.notificationhubs.models.PolicyKeyType;

/**
 * Samples for Namespaces RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/AuthorizationRuleRegenerateKey.json
     */
    /**
     * Sample code: Namespaces_RegenerateKeys.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        namespacesRegenerateKeys(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().regenerateKeysWithResponse("5ktrial", "nh-sdk-ns", "RootManageSharedAccessKey",
            new PolicyKeyResource().withPolicyKey(PolicyKeyType.PRIMARY_KEY), com.azure.core.util.Context.NONE);
    }
}
