
/**
 * Samples for WebApps ListBackups.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListWebAppBackups.json
     */
    /**
     * Sample code: List web app backups.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listWebAppBackups(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listBackups("testrg123", "sitef6141", com.azure.core.util.Context.NONE);
    }
}
