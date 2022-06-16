import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.AccessIdName;

/** Samples for TenantAccessGit RegeneratePrimaryKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementTenantAccessRegenerateKey.json
     */
    /**
     * Sample code: ApiManagementTenantAccessRegenerateKey.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementTenantAccessRegenerateKey(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tenantAccessGits()
            .regeneratePrimaryKeyWithResponse("rg1", "apimService1", AccessIdName.ACCESS, Context.NONE);
    }
}
