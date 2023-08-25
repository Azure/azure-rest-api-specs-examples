/** Samples for ProductApi CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateProductApi.json
     */
    /**
     * Sample code: ApiManagementCreateProductApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateProductApi(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productApis()
            .createOrUpdateWithResponse(
                "rg1", "apimService1", "testproduct", "echo-api", com.azure.core.util.Context.NONE);
    }
}
