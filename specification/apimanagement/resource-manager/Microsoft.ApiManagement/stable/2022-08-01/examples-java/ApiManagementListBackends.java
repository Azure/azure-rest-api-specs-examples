/** Samples for Backend ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListBackends.json
     */
    /**
     * Sample code: ApiManagementListBackends.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListBackends(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.backends().listByService("rg1", "apimService1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
