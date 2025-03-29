
/**
 * Samples for PolicySetDefinitionVersions ListAllBuiltins.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * listAllBuiltInPolicySetDefinitionVersions.json
     */
    /**
     * Sample code: List all built-in policy definition versions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllBuiltInPolicyDefinitionVersions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitionVersions()
            .listAllBuiltinsWithResponse(com.azure.core.util.Context.NONE);
    }
}
