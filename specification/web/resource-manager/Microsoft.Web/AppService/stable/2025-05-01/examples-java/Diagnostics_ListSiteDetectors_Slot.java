
/**
 * Samples for Diagnostics ListSiteDetectorsSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ListSiteDetectors_Slot.json
     */
    /**
     * Sample code: List App Detectors.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAppDetectors(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().listSiteDetectorsSlot("Sample-WestUSResourceGroup", "SampleApp",
            "availability", "Production", com.azure.core.util.Context.NONE);
    }
}
