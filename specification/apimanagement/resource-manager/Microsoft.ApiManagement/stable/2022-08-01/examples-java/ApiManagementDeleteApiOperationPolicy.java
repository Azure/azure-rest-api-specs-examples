import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for ApiOperationPolicy Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteApiOperationPolicy.json
     */
    /**
     * Sample code: ApiManagementDeleteApiOperationPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteApiOperationPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiOperationPolicies()
            .deleteWithResponse(
                "rg1",
                "apimService1",
                "testapi",
                "testoperation",
                PolicyIdName.POLICY,
                "*",
                com.azure.core.util.Context.NONE);
    }
}
