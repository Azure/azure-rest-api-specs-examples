
/**
 * Samples for PolicySetDefinitions DeleteAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * deletePolicySetDefinitionAtManagementGroup.json
     */
    /**
     * Sample code: Delete a policy set definition at management group level.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deleteAPolicySetDefinitionAtManagementGroupLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitions().deleteAtManagementGroupWithResponse(
            "MyManagementGroup", "CostManagement", com.azure.core.util.Context.NONE);
    }
}
