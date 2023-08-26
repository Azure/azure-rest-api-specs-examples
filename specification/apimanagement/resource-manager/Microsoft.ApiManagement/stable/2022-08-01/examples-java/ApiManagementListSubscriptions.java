/** Samples for Subscription List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListSubscriptions.json
     */
    /**
     * Sample code: ApiManagementListSubscriptions.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListSubscriptions(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.subscriptions().list("rg1", "apimService1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
