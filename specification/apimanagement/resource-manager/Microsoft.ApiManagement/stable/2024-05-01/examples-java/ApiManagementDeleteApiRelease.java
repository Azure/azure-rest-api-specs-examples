
/**
 * Samples for ApiRelease Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteApiRelease.json
     */
    /**
     * Sample code: ApiManagementDeleteApiRelease.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteApiRelease(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiReleases().deleteWithResponse("rg1", "apimService1", "5a5fcc09124a7fa9b89f2f1d", "testrev", "*",
            com.azure.core.util.Context.NONE);
    }
}
