
/**
 * Samples for AppServiceEnvironments GetAseV3NetworkingConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_GetAseV3NetworkingConfiguration.json
     */
    /**
     * Sample code: Get networking configuration of an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getNetworkingConfigurationOfAnAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().getAseV3NetworkingConfigurationWithResponse("test-rg",
            "test-ase", com.azure.core.util.Context.NONE);
    }
}
