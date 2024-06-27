
/**
 * Samples for PolicyDefinitions ListByManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/
     * listPolicyDefinitionsByManagementGroup.json
     */
    /**
     * Sample code: List policy definitions by management group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPolicyDefinitionsByManagementGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyDefinitions()
            .listByManagementGroup("MyManagementGroup", null, null, com.azure.core.util.Context.NONE);
    }
}
