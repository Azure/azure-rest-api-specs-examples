
/**
 * Samples for DeploymentSettings ListByClusters.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * ListDeploymentSettingsByCluster.json
     */
    /**
     * Sample code: List Deployment Settings.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listDeploymentSettings(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.deploymentSettings().listByClusters("test-rg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
