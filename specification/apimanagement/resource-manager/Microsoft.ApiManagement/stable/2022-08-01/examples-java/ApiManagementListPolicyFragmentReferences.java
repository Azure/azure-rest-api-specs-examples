/** Samples for PolicyFragment ListReferences. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListPolicyFragmentReferences.json
     */
    /**
     * Sample code: ApiManagementListPolicyFragmentReferences.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListPolicyFragmentReferences(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .policyFragments()
            .listReferencesWithResponse(
                "rg1", "apimService1", "policyFragment1", null, null, com.azure.core.util.Context.NONE);
    }
}
