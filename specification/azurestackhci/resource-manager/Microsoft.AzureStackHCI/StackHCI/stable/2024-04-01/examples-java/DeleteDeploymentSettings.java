
/**
 * Samples for DeploymentSettings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * DeleteDeploymentSettings.json
     */
    /**
     * Sample code: Delete Deployment Settings.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteDeploymentSettings(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.deploymentSettings().delete("test-rg", "myCluster", "default", com.azure.core.util.Context.NONE);
    }
}
