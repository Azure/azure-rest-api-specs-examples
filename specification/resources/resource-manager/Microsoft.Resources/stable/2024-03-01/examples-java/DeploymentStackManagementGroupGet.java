
/**
 * Samples for DeploymentStacks GetAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * DeploymentStackManagementGroupGet.json
     */
    /**
     * Sample code: DeploymentStacksManagementGroupGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentStacksManagementGroupGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().deploymentStackClient().getDeploymentStacks()
            .getAtManagementGroupWithResponse("myMg", "simpleDeploymentStack", com.azure.core.util.Context.NONE);
    }
}
