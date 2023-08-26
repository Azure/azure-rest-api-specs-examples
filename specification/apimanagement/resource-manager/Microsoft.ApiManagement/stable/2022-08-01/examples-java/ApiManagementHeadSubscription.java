/** Samples for Subscription GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadSubscription.json
     */
    /**
     * Sample code: ApiManagementHeadSubscription.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadSubscription(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .subscriptions()
            .getEntityTagWithResponse(
                "rg1", "apimService1", "5931a769d8d14f0ad8ce13b8", com.azure.core.util.Context.NONE);
    }
}
