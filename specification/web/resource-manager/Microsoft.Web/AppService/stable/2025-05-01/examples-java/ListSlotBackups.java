
/**
 * Samples for WebApps ListSiteBackups.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListSlotBackups.json
     */
    /**
     * Sample code: List backups.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listBackups(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listSiteBackups("testrg123", "tests346", com.azure.core.util.Context.NONE);
    }
}
