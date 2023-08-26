/** Samples for ApiRelease ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListApiReleases.json
     */
    /**
     * Sample code: ApiManagementListApiReleases.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApiReleases(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiReleases()
            .listByService("rg1", "apimService1", "a1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
