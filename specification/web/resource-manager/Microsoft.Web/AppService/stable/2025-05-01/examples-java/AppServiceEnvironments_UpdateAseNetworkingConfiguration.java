
import com.azure.resourcemanager.appservice.fluent.models.AseV3NetworkingConfigurationInner;

/**
 * Samples for AppServiceEnvironments UpdateAseNetworkingConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_UpdateAseNetworkingConfiguration.json
     */
    /**
     * Sample code: Update networking configuration of an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void updateNetworkingConfigurationOfAnAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().updateAseNetworkingConfigurationWithResponse("test-rg",
            "test-ase", new AseV3NetworkingConfigurationInner().withFtpEnabled(true).withRemoteDebugEnabled(true),
            com.azure.core.util.Context.NONE);
    }
}
