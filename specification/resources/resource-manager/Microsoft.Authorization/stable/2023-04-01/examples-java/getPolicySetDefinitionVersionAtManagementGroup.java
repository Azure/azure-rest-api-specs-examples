
/**
 * Samples for PolicySetDefinitionVersions GetAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * getPolicySetDefinitionVersionAtManagementGroup.json
     */
    /**
     * Sample code: Retrieve a policy set definition version at management group level.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveAPolicySetDefinitionVersionAtManagementGroupLevel(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitionVersions()
            .getAtManagementGroupWithResponse("MyManagementGroup", "CostManagement", "1.2.1", null,
                com.azure.core.util.Context.NONE);
    }
}
