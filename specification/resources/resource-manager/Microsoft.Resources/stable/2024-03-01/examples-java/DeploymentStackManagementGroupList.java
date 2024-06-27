
/**
 * Samples for DeploymentStacks ListAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * DeploymentStackManagementGroupList.json
     */
    /**
     * Sample code: DeploymentStacksManagementGroupList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentStacksManagementGroupList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().deploymentStackClient().getDeploymentStacks().listAtManagementGroup("myMg",
            com.azure.core.util.Context.NONE);
    }
}
