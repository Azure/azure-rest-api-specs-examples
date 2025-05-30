
import com.azure.resourcemanager.apimanagement.models.AccessIdName;

/**
 * Samples for TenantAccess GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadTenantAccess.json
     */
    /**
     * Sample code: ApiManagementHeadTenantAccess.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadTenantAccess(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tenantAccess().getEntityTagWithResponse("rg1", "apimService1", AccessIdName.ACCESS,
            com.azure.core.util.Context.NONE);
    }
}
