import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for GraphQLApiResolverPolicy Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteGraphQLApiResolverPolicy.json
     */
    /**
     * Sample code: ApiManagementDeleteGraphQLApiResolverPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteGraphQLApiResolverPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .graphQLApiResolverPolicies()
            .deleteWithResponse(
                "rg1",
                "apimService1",
                "testapi",
                "testResolver",
                PolicyIdName.POLICY,
                "*",
                com.azure.core.util.Context.NONE);
    }
}
