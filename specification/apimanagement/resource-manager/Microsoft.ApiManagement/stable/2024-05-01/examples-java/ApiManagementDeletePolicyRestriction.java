
/**
 * Samples for PolicyRestriction Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeletePolicyRestriction.json
     */
    /**
     * Sample code: ApiManagementDeletePolicyRestriction.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeletePolicyRestriction(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.policyRestrictions().deleteWithResponse("rg1", "apimService1", "policyRestriction1", "*",
            com.azure.core.util.Context.NONE);
    }
}
