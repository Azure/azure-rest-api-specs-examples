/** Samples for PolicyFragment Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetPolicyFragment.json
     */
    /**
     * Sample code: ApiManagementGetPolicyFragment.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetPolicyFragment(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .policyFragments()
            .getWithResponse("rg1", "apimService1", "policyFragment1", null, com.azure.core.util.Context.NONE);
    }
}
