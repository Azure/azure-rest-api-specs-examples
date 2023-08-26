/** Samples for Cache ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListCaches.json
     */
    /**
     * Sample code: ApiManagementListCaches.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListCaches(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.caches().listByService("rg1", "apimService1", null, null, com.azure.core.util.Context.NONE);
    }
}
