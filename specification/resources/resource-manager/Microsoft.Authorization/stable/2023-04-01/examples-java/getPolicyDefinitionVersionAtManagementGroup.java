
/**
 * Samples for PolicyDefinitionVersions GetAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * getPolicyDefinitionVersionAtManagementGroup.json
     */
    /**
     * Sample code: Retrieve a policy definition version at management group level.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        retrieveAPolicyDefinitionVersionAtManagementGroupLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyDefinitionVersions()
            .getAtManagementGroupWithResponse("MyManagementGroup", "ResourceNaming", "1.2.1",
                com.azure.core.util.Context.NONE);
    }
}
