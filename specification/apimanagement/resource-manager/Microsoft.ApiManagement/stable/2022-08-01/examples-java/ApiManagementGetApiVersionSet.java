/** Samples for ApiVersionSet Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiVersionSet.json
     */
    /**
     * Sample code: ApiManagementGetApiVersionSet.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiVersionSet(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiVersionSets().getWithResponse("rg1", "apimService1", "vs1", com.azure.core.util.Context.NONE);
    }
}
