
/**
 * Samples for PolicySetDefinitionVersions GetBuiltIn.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * getBuiltInPolicySetDefinitionVersion.json
     */
    /**
     * Sample code: Retrieve a built-in policy set definition version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        retrieveABuiltInPolicySetDefinitionVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitionVersions().getBuiltInWithResponse(
            "1f3afdf9-d0c9-4c3d-847f-89da613e70a8", "1.2.1", null, com.azure.core.util.Context.NONE);
    }
}
