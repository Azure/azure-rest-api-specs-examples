
/**
 * Samples for PolicyDefinitionVersions GetBuiltIn.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * getBuiltinPolicyDefinitionVersion.json
     */
    /**
     * Sample code: Retrieve a built-in policy definition version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveABuiltInPolicyDefinitionVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyDefinitionVersions()
            .getBuiltInWithResponse("7433c107-6db4-4ad1-b57a-a76dce0154a1", "1.2.1", com.azure.core.util.Context.NONE);
    }
}
