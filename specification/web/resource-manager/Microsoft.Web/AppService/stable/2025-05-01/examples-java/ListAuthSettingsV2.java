
/**
 * Samples for WebApps GetAuthSettingsV2.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListAuthSettingsV2.json
     */
    /**
     * Sample code: List Auth Settings V2.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAuthSettingsV2(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getAuthSettingsV2WithResponse("testrg123", "sitef6141",
            com.azure.core.util.Context.NONE);
    }
}
