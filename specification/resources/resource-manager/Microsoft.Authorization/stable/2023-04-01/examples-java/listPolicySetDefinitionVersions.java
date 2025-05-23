
/**
 * Samples for PolicySetDefinitionVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * listPolicySetDefinitionVersions.json
     */
    /**
     * Sample code: List policy set definitions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPolicySetDefinitions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitionVersions().list("CostManagement", null,
            null, com.azure.core.util.Context.NONE);
    }
}
