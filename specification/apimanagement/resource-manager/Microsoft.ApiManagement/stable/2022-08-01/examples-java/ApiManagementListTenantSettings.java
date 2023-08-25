/** Samples for TenantSettings ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListTenantSettings.json
     */
    /**
     * Sample code: ApiManagementListTenantSettings.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListTenantSettings(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tenantSettings().listByService("rg1", "apimService1", null, com.azure.core.util.Context.NONE);
    }
}
