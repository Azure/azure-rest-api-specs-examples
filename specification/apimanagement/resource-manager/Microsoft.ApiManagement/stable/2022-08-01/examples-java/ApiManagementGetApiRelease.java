/** Samples for ApiRelease Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiRelease.json
     */
    /**
     * Sample code: ApiManagementGetApiRelease.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiRelease(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiReleases()
            .getWithResponse("rg1", "apimService1", "a1", "5a7cb545298324c53224a799", com.azure.core.util.Context.NONE);
    }
}
