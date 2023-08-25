/** Samples for ProductSubscriptions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListProductSubscriptions.json
     */
    /**
     * Sample code: ApiManagementListProductSubscriptions.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListProductSubscriptions(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productSubscriptions()
            .list(
                "rg1", "apimService1", "5600b57e7e8880006a060002", null, null, null, com.azure.core.util.Context.NONE);
    }
}
