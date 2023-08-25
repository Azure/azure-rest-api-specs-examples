/** Samples for Api Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteApi.json
     */
    /**
     * Sample code: ApiManagementDeleteApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteApi(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .deleteWithResponse("rg1", "apimService1", "echo-api", "*", null, com.azure.core.util.Context.NONE);
    }
}
