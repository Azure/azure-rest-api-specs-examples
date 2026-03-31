
/**
 * Samples for WebApps GetSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetWebAppSlot.json
     */
    /**
     * Sample code: Get Web App Slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getWebAppSlot(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getSlotWithResponse("testrg123", "sitef6141", "staging",
            com.azure.core.util.Context.NONE);
    }
}
