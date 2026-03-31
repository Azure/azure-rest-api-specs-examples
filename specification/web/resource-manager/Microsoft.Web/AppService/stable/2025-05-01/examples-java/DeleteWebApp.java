
/**
 * Samples for WebApps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteWebApp.json
     */
    /**
     * Sample code: Delete Web app.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deleteWebApp(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().deleteWithResponse("testrg123", "sitef6141", null, null,
            com.azure.core.util.Context.NONE);
    }
}
