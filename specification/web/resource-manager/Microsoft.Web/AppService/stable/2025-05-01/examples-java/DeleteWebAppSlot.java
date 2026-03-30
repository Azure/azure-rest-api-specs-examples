
/**
 * Samples for WebApps DeleteSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteWebAppSlot.json
     */
    /**
     * Sample code: Delete Web App Slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deleteWebAppSlot(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().deleteSlotWithResponse("testrg123", "sitef6141", "staging", null, null,
            com.azure.core.util.Context.NONE);
    }
}
