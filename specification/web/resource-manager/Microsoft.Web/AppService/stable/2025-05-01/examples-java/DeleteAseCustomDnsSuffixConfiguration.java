
/**
 * Samples for AppServiceEnvironments DeleteAseCustomDnsSuffixConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteAseCustomDnsSuffixConfiguration.json
     */
    /**
     * Sample code: Delete ASE custom DNS suffix configuration.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        deleteASECustomDNSSuffixConfiguration(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().deleteAseCustomDnsSuffixConfigurationWithResponse("test-rg",
            "test-ase", com.azure.core.util.Context.NONE);
    }
}
