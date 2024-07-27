
/**
 * Samples for PolicyDefinitions DeleteAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/
     * deletePolicyDefinitionAtManagementGroup.json
     */
    /**
     * Sample code: Delete a policy definition at management group level.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deleteAPolicyDefinitionAtManagementGroupLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyDefinitions().deleteAtManagementGroupWithResponse(
            "ResourceNaming", "MyManagementGroup", com.azure.core.util.Context.NONE);
    }
}
