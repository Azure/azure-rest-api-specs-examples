/** Samples for GraphQLApiResolver CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGraphQLApiResolver.json
     */
    /**
     * Sample code: ApiManagementCreateGraphQLApiResolver.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateGraphQLApiResolver(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .graphQLApiResolvers()
            .define("newResolver")
            .withExistingApi("rg1", "apimService1", "someAPI")
            .withDisplayName("Query Users")
            .withPath("Query/users")
            .withDescription("A GraphQL Resolver example")
            .create();
    }
}
