
/**
 * Samples for PolicySetDefinitionVersions ListAll.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * listAllPolicySetDefinitionVersions.json
     */
    /**
     * Sample code: List all policy definition versions at subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAllPolicyDefinitionVersionsAtSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitionVersions()
            .listAllWithResponse(com.azure.core.util.Context.NONE);
    }
}
