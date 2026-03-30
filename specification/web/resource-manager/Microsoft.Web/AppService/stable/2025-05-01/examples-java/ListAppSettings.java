
/**
 * Samples for WebApps ListApplicationSettings.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListAppSettings.json
     */
    /**
     * Sample code: List App Settings.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAppSettings(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listApplicationSettingsWithResponse("testrg123", "sitef6141",
            com.azure.core.util.Context.NONE);
    }
}
