
/**
 * Samples for PolicySetDefinitionVersions ListByManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * listPolicySetDefinitionVersionsByManagementGroup.json
     */
    /**
     * Sample code: List policy set definitions at management group level.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listPolicySetDefinitionsAtManagementGroupLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitionVersions()
            .listByManagementGroup("MyManagementGroup", "CostManagement", null, null, com.azure.core.util.Context.NONE);
    }
}
