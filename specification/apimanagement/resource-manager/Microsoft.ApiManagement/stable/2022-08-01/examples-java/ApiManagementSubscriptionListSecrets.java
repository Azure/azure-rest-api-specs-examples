/** Samples for Subscription ListSecrets. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementSubscriptionListSecrets.json
     */
    /**
     * Sample code: ApiManagementSubscriptionListSecrets.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementSubscriptionListSecrets(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .subscriptions()
            .listSecretsWithResponse(
                "rg1", "apimService1", "5931a769d8d14f0ad8ce13b8", com.azure.core.util.Context.NONE);
    }
}
