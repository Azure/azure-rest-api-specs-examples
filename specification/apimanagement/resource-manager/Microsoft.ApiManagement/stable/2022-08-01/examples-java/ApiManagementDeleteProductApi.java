/** Samples for ProductApi Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteProductApi.json
     */
    /**
     * Sample code: ApiManagementDeleteProductApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteProductApi(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productApis()
            .deleteWithResponse("rg1", "apimService1", "testproduct", "echo-api", com.azure.core.util.Context.NONE);
    }
}
