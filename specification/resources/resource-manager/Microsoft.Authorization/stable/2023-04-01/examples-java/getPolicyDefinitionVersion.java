
/**
 * Samples for PolicyDefinitionVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * getPolicyDefinitionVersion.json
     */
    /**
     * Sample code: Retrieve a policy definition version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveAPolicyDefinitionVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyDefinitionVersions()
            .getWithResponse("ResourceNaming", "1.2.1", com.azure.core.util.Context.NONE);
    }
}
