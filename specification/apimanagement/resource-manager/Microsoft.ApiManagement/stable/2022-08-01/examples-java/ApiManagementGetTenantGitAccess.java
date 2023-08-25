import com.azure.resourcemanager.apimanagement.models.AccessIdName;

/** Samples for TenantAccess Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetTenantGitAccess.json
     */
    /**
     * Sample code: ApiManagementGetTenantGitAccess.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetTenantGitAccess(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tenantAccess()
            .getWithResponse("rg1", "apimService1", AccessIdName.GIT_ACCESS, com.azure.core.util.Context.NONE);
    }
}
