
/**
 * Samples for PolicyDefinitionVersions ListBuiltIn.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * listBuiltInPolicyDefinitionVersions.json
     */
    /**
     * Sample code: List built-in policy definition versions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listBuiltInPolicyDefinitionVersions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyDefinitionVersions()
            .listBuiltIn("06a78e20-9358-41c9-923c-fb736d382a12", null, com.azure.core.util.Context.NONE);
    }
}
