
/** Samples for PolicyDefinitions GetAtManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/
     * getPolicyDefinitionAtManagementGroup.json
     */
    /**
     * Sample code: Retrieve a policy definition at management group level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        retrieveAPolicyDefinitionAtManagementGroupLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyDefinitions()
            .getAtManagementGroupWithResponse("ResourceNaming", "MyManagementGroup", com.azure.core.util.Context.NONE);
    }
}
