import com.azure.resourcemanager.apimanagement.models.PolicyFragmentContentFormat;

/** Samples for PolicyFragment Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetPolicyFragmentFormat.json
     */
    /**
     * Sample code: ApiManagementGetPolicyFragmentFormat.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetPolicyFragmentFormat(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .policyFragments()
            .getWithResponse(
                "rg1",
                "apimService1",
                "policyFragment1",
                PolicyFragmentContentFormat.RAWXML,
                com.azure.core.util.Context.NONE);
    }
}
