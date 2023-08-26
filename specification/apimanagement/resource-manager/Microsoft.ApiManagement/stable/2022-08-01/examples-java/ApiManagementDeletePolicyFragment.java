/** Samples for PolicyFragment Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeletePolicyFragment.json
     */
    /**
     * Sample code: ApiManagementDeletePolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeletePolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .policyFragments()
            .deleteWithResponse("rg1", "apimService1", "policyFragment1", "*", com.azure.core.util.Context.NONE);
    }
}
