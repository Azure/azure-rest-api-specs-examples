/** Samples for GraphQLApiResolver GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadGraphQLApiResolver.json
     */
    /**
     * Sample code: ApiManagementHeadGraphQLApiResolver.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadGraphQLApiResolver(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .graphQLApiResolvers()
            .getEntityTagWithResponse(
                "rg1",
                "apimService1",
                "57d2ef278aa04f0888cba3f3",
                "57d2ef278aa04f0ad01d6cdc",
                com.azure.core.util.Context.NONE);
    }
}
