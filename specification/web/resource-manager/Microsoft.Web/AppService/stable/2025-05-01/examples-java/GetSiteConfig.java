
/**
 * Samples for WebApps GetConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetSiteConfig.json
     */
    /**
     * Sample code: Get Site Config.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getSiteConfig(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getConfigurationWithResponse("testrg123", "sitef6141",
            com.azure.core.util.Context.NONE);
    }
}
