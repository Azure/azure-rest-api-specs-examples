
/**
 * Samples for PolicyDefinitionVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * deletePolicyDefinitionVersion.json
     */
    /**
     * Sample code: Delete a policy definition version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAPolicyDefinitionVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyDefinitionVersions()
            .deleteWithResponse("ResourceNaming", "1.2.1", com.azure.core.util.Context.NONE);
    }
}
