
/**
 * Samples for PolicySetDefinitionVersions ListBuiltIn.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * listBuiltInPolicySetDefinitionVersions.json
     */
    /**
     * Sample code: List built-in policy set definitions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listBuiltInPolicySetDefinitions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitionVersions()
            .listBuiltIn("1f3afdf9-d0c9-4c3d-847f-89da613e70a8", null, null, com.azure.core.util.Context.NONE);
    }
}
