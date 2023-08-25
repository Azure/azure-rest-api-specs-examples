import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for ApiPolicy GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadApiPolicy.json
     */
    /**
     * Sample code: ApiManagementHeadApiPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadApiPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiPolicies()
            .getEntityTagWithResponse(
                "rg1",
                "apimService1",
                "57d1f7558aa04f15146d9d8a",
                PolicyIdName.POLICY,
                com.azure.core.util.Context.NONE);
    }
}
