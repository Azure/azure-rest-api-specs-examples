
/** Samples for PolicySetDefinitions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/
     * getPolicySetDefinition.json
     */
    /**
     * Sample code: Retrieve a policy set definition.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveAPolicySetDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitions().getWithResponse("CostManagement",
            com.azure.core.util.Context.NONE);
    }
}
