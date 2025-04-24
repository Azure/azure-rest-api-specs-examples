
/**
 * Samples for PolicyRestriction GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadPolicyRestriction.json
     */
    /**
     * Sample code: ApiManagementHeadPolicyRestriction.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadPolicyRestriction(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.policyRestrictions().getEntityTagWithResponse("rg1", "apimService1", "policyRestriction1",
            com.azure.core.util.Context.NONE);
    }
}
