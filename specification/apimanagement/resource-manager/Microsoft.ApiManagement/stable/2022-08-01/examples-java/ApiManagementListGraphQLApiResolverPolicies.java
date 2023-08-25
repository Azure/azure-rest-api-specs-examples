/** Samples for GraphQLApiResolverPolicy ListByResolver. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGraphQLApiResolverPolicies.json
     */
    /**
     * Sample code: ApiManagementListGraphQLApiResolverPolicies.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListGraphQLApiResolverPolicies(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .graphQLApiResolverPolicies()
            .listByResolver(
                "rg1",
                "apimService1",
                "599e2953193c3c0bd0b3e2fa",
                "599e29ab193c3c0bd0b3e2fb",
                com.azure.core.util.Context.NONE);
    }
}
