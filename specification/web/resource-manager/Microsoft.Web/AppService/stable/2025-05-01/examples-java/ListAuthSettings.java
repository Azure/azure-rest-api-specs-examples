
/**
 * Samples for WebApps GetAuthSettings.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListAuthSettings.json
     */
    /**
     * Sample code: List Auth Settings.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAuthSettings(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getAuthSettingsWithResponse("testrg123", "sitef6141",
            com.azure.core.util.Context.NONE);
    }
}
