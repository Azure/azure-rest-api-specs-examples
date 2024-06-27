
/**
 * Samples for DeploymentStacks DeleteAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * DeploymentStackManagementGroupDelete.json
     */
    /**
     * Sample code: DeploymentStacksManagementGroupDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentStacksManagementGroupDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().deploymentStackClient().getDeploymentStacks().deleteAtManagementGroup("myMg",
            "simpleDeploymentStack", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
