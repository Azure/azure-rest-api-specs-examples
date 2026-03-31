
/**
 * Samples for Diagnostics ListSiteDetectorsSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ListSiteDetectorsSlot_Slot.json
     */
    /**
     * Sample code: List App Slot Detectors.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAppSlotDetectors(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().listSiteDetectorsSlot("Sample-WestUSResourceGroup", "SampleApp",
            "availability", "staging", com.azure.core.util.Context.NONE);
    }
}
