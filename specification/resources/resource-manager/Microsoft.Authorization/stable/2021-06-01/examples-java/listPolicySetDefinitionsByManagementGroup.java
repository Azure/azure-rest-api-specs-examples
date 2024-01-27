
/** Samples for PolicySetDefinitions ListByManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/
     * listPolicySetDefinitionsByManagementGroup.json
     */
    /**
     * Sample code: List policy set definitions at management group level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listPolicySetDefinitionsAtManagementGroupLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitions()
            .listByManagementGroup("MyManagementGroup", null, null, com.azure.core.util.Context.NONE);
    }
}
