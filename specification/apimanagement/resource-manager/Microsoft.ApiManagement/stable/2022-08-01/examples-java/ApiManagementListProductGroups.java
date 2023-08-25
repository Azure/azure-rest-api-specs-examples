/** Samples for ProductGroup ListByProduct. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListProductGroups.json
     */
    /**
     * Sample code: ApiManagementListProductGroups.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListProductGroups(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productGroups()
            .listByProduct(
                "rg1", "apimService1", "5600b57e7e8880006a060002", null, null, null, com.azure.core.util.Context.NONE);
    }
}
