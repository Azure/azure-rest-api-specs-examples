
/**
 * Samples for DeploymentStacks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * DeploymentStackResourceGroupDelete.json
     */
    /**
     * Sample code: DeploymentStacksResourceGroupDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentStacksResourceGroupDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().deploymentStackClient().getDeploymentStacks().delete("deploymentStacksRG",
            "simpleDeploymentStack", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
