import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for ApiOperationPolicy GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadApiOperationPolicy.json
     */
    /**
     * Sample code: ApiManagementHeadApiOperationPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadApiOperationPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiOperationPolicies()
            .getEntityTagWithResponse(
                "rg1",
                "apimService1",
                "5600b539c53f5b0062040001",
                "5600b53ac53f5b0062080006",
                PolicyIdName.POLICY,
                com.azure.core.util.Context.NONE);
    }
}
