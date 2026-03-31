
/**
 * Samples for WebApps GetInstanceInfoSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetSiteInstanceInfo_Slot.json
     */
    /**
     * Sample code: Get site instance info.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getSiteInstanceInfo(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getInstanceInfoSlotWithResponse("testrg123", "tests346", "134987120",
            "staging", com.azure.core.util.Context.NONE);
    }
}
