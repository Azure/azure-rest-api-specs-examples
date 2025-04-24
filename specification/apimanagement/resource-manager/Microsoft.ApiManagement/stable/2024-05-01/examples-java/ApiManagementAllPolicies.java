
/**
 * Samples for AllPolicies ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementAllPolicies.json
     */
    /**
     * Sample code: ApiManagementListPolicyRestrictions.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListPolicyRestrictions(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.allPolicies().listByService("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
