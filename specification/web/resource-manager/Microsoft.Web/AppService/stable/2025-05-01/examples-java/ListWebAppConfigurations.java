
/**
 * Samples for WebApps ListConfigurations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListWebAppConfigurations.json
     */
    /**
     * Sample code: List web app configurations.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listWebAppConfigurations(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listConfigurations("testrg123", "sitef6141",
            com.azure.core.util.Context.NONE);
    }
}
