/** Samples for ProductApi ListByProduct. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListProductApis.json
     */
    /**
     * Sample code: ApiManagementListProductApis.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListProductApis(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productApis()
            .listByProduct(
                "rg1", "apimService1", "5768181ea40f7eb6c49f6ac7", null, null, null, com.azure.core.util.Context.NONE);
    }
}
