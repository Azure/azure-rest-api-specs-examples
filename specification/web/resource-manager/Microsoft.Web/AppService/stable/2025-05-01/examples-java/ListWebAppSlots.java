
/**
 * Samples for WebApps ListSlots.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListWebAppSlots.json
     */
    /**
     * Sample code: List Web App Slots.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listWebAppSlots(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listSlots("testrg123", "sitef6141", com.azure.core.util.Context.NONE);
    }
}
