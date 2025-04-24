
import com.azure.resourcemanager.apimanagement.models.ResolverContract;

/**
 * Samples for GraphQLApiResolver Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateGraphQLApiResolver.json
     */
    /**
     * Sample code: ApiManagementUpdateGraphQLApiResolver.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateGraphQLApiResolver(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        ResolverContract resource = manager.graphQLApiResolvers()
            .getWithResponse("rg1", "apimService1", "echo-api", "resolverId", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withDisplayName("Query AdminUsers").withPath("Query/adminUsers")
            .withDescription("A GraphQL Resolver example").withIfMatch("*").apply();
    }
}
