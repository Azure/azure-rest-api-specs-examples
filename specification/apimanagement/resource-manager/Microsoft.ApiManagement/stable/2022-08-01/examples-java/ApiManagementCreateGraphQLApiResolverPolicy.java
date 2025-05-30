
import com.azure.resourcemanager.apimanagement.fluent.models.PolicyContractInner;
import com.azure.resourcemanager.apimanagement.models.PolicyContentFormat;
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for GraphQLApiResolverPolicy CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/
     * ApiManagementCreateGraphQLApiResolverPolicy.json
     */
    /**
     * Sample code: ApiManagementCreateGraphQLApiResolverPolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateGraphQLApiResolverPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.graphQLApiResolverPolicies().createOrUpdateWithResponse("rg1", "apimService1",
            "5600b57e7e8880006a040001", "5600b57e7e8880006a080001", PolicyIdName.POLICY,
            new PolicyContractInner().withValue(
                "<http-data-source><http-request><set-method>GET</set-method><set-backend-service base-url=\"https://some.service.com\" /><set-url>/api/users</set-url></http-request></http-data-source>")
                .withFormat(PolicyContentFormat.XML),
            "*", com.azure.core.util.Context.NONE);
    }
}
