
/**
 * Samples for PolicyDefinitionVersions DeleteAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * deletePolicyDefinitionVersionAtManagementGroup.json
     */
    /**
     * Sample code: Delete a policy definition version at management group level.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deleteAPolicyDefinitionVersionAtManagementGroupLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyDefinitionVersions()
            .deleteAtManagementGroupWithResponse("MyManagementGroup", "ResourceNaming", "1.2.1",
                com.azure.core.util.Context.NONE);
    }
}
