/** Samples for ApiVersionSet Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteApiVersionSet.json
     */
    /**
     * Sample code: ApiManagementDeleteApiVersionSet.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteApiVersionSet(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiVersionSets().deleteWithResponse("rg1", "apimService1", "a1", "*", com.azure.core.util.Context.NONE);
    }
}
