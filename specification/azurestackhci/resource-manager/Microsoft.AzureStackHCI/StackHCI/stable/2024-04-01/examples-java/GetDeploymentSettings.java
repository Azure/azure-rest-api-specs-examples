
/**
 * Samples for DeploymentSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * GetDeploymentSettings.json
     */
    /**
     * Sample code: Get Deployment Settings.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getDeploymentSettings(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.deploymentSettings().getWithResponse("test-rg", "myCluster", "default",
            com.azure.core.util.Context.NONE);
    }
}
