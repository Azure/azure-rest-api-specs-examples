/** Samples for ApiVersionSet ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListApiVersionSets.json
     */
    /**
     * Sample code: ApiManagementListApiVersionSets.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApiVersionSets(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiVersionSets()
            .listByService("rg1", "apimService1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
