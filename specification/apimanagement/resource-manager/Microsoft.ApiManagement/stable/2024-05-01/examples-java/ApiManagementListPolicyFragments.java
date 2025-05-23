
/**
 * Samples for PolicyFragment ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListPolicyFragments.json
     */
    /**
     * Sample code: ApiManagementListPolicyFragments.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListPolicyFragments(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.policyFragments().listByService("rg1", "apimService1", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
