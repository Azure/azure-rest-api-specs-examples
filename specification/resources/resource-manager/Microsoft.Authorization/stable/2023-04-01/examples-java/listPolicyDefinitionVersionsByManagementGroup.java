
/**
 * Samples for PolicyDefinitionVersions ListByManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * listPolicyDefinitionVersionsByManagementGroup.json
     */
    /**
     * Sample code: List policy definition versions by management group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listPolicyDefinitionVersionsByManagementGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyDefinitionVersions()
            .listByManagementGroup("MyManagementGroup", "ResourceNaming", null, com.azure.core.util.Context.NONE);
    }
}
