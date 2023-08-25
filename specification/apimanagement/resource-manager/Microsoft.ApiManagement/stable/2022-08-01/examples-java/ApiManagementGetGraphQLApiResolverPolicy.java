import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for GraphQLApiResolverPolicy Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetGraphQLApiResolverPolicy.json
     */
    /**
     * Sample code: ApiManagementGetGraphQLApiResolverPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetGraphQLApiResolverPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .graphQLApiResolverPolicies()
            .getWithResponse(
                "rg1",
                "apimService1",
                "5600b539c53f5b0062040001",
                "5600b53ac53f5b0062080006",
                PolicyIdName.POLICY,
                null,
                com.azure.core.util.Context.NONE);
    }
}
