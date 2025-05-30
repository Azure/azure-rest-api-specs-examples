
/**
 * Samples for PolicySetDefinitionVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * getPolicySetDefinitionVersion.json
     */
    /**
     * Sample code: Retrieve a policy set definition version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveAPolicySetDefinitionVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitionVersions()
            .getWithResponse("CostManagement", "1.2.1", null, com.azure.core.util.Context.NONE);
    }
}
