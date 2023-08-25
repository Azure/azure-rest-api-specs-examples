/** Samples for PolicyFragment GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadPolicyFragment.json
     */
    /**
     * Sample code: ApiManagementHeadPolicyFragment.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadPolicyFragment(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .policyFragments()
            .getEntityTagWithResponse("rg1", "apimService1", "policyFragment1", com.azure.core.util.Context.NONE);
    }
}
