
/**
 * Samples for DeploymentSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/GetDeploymentSettings.json
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
