
/**
 * Samples for GraphQLApiResolver Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetGraphQLApiResolver.json
     */
    /**
     * Sample code: ApiManagementGetGraphQLApiResolver.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetGraphQLApiResolver(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.graphQLApiResolvers().getWithResponse("rg1", "apimService1", "57d2ef278aa04f0888cba3f3",
            "57d2ef278aa04f0ad01d6cdc", com.azure.core.util.Context.NONE);
    }
}
