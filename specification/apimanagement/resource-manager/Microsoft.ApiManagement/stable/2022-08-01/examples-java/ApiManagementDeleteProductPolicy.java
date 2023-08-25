import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for ProductPolicy Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteProductPolicy.json
     */
    /**
     * Sample code: ApiManagementDeleteProductPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteProductPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productPolicies()
            .deleteWithResponse(
                "rg1", "apimService1", "testproduct", PolicyIdName.POLICY, "*", com.azure.core.util.Context.NONE);
    }
}
