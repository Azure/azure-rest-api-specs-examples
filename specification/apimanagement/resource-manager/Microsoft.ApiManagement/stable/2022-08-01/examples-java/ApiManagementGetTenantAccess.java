import com.azure.resourcemanager.apimanagement.models.AccessIdName;

/** Samples for TenantAccess Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetTenantAccess.json
     */
    /**
     * Sample code: ApiManagementGetTenantAccess.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetTenantAccess(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tenantAccess()
            .getWithResponse("rg1", "apimService1", AccessIdName.ACCESS, com.azure.core.util.Context.NONE);
    }
}
