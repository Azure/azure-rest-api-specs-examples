import com.azure.core.util.Context;

/** Samples for PolicySetDefinitions GetBuiltIn. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getBuiltInPolicySetDefinition.json
     */
    /**
     * Sample code: Retrieve a built-in policy set definition.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveABuiltInPolicySetDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicySetDefinitions()
            .getBuiltInWithResponse("1f3afdf9-d0c9-4c3d-847f-89da613e70a8", Context.NONE);
    }
}
