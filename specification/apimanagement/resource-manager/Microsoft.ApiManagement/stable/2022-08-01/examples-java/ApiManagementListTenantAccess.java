/** Samples for TenantAccess ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListTenantAccess.json
     */
    /**
     * Sample code: ApiManagementListTenantAccess.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListTenantAccess(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tenantAccess().listByService("rg1", "apimService1", null, com.azure.core.util.Context.NONE);
    }
}
