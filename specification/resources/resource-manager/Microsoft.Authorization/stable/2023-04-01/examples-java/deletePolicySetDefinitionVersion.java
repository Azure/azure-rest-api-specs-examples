
/**
 * Samples for PolicySetDefinitionVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * deletePolicySetDefinitionVersion.json
     */
    /**
     * Sample code: Delete a policy set definition version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAPolicySetDefinitionVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitionVersions()
            .deleteWithResponse("CostManagement", "1.2.1", com.azure.core.util.Context.NONE);
    }
}
