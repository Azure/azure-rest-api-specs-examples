/** Samples for UserSubscription List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListUserSubscriptions.json
     */
    /**
     * Sample code: ApiManagementListUserSubscriptions.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListUserSubscriptions(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .userSubscriptions()
            .list(
                "rg1", "apimService1", "57681833a40f7eb6c49f6acf", null, null, null, com.azure.core.util.Context.NONE);
    }
}
