
/**
 * Samples for PolicySetDefinitionVersions ListAllAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * listAllPolicySetDefinitionVersionsByManagementGroup.json
     */
    /**
     * Sample code: List all policy definition versions at management group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAllPolicyDefinitionVersionsAtManagementGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitionVersions()
            .listAllAtManagementGroupWithResponse("MyManagementGroup", com.azure.core.util.Context.NONE);
    }
}
