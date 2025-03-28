
/**
 * Samples for PolicyDefinitions GetAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
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
            .getAtManagementGroupWithResponse("MyManagementGroup", "ResourceNaming", com.azure.core.util.Context.NONE);
    }
}
