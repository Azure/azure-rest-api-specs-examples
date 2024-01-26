
/** Samples for PolicySetDefinitions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/
     * deletePolicySetDefinition.json
     */
    /**
     * Sample code: Delete a policy set definition.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAPolicySetDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitions().deleteWithResponse("CostManagement",
            com.azure.core.util.Context.NONE);
    }
}
