/** Samples for GraphQLApiResolver ListByApi. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGraphQLApiResolvers.json
     */
    /**
     * Sample code: ApiManagementListGraphQLApiResolvers.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListGraphQLApiResolvers(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .graphQLApiResolvers()
            .listByApi(
                "rg1", "apimService1", "57d2ef278aa04f0888cba3f3", null, null, null, com.azure.core.util.Context.NONE);
    }
}
