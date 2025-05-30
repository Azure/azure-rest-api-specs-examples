
/**
 * Samples for Policy ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListPolicies.json
     */
    /**
     * Sample code: ApiManagementListPolicies.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListPolicies(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.policies().listByService("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
