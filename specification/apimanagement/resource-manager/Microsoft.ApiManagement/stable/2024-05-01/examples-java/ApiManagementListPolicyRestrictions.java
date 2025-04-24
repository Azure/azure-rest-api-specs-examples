
/**
 * Samples for PolicyRestriction ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListPolicyRestrictions.json
     */
    /**
     * Sample code: ApiManagementListPolicyRestrictions.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListPolicyRestrictions(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.policyRestrictions().listByService("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
