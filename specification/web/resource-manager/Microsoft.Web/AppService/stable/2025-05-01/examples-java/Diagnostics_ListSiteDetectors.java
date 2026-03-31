
/**
 * Samples for Diagnostics ListSiteDetectors.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ListSiteDetectors.json
     */
    /**
     * Sample code: List App Detectors.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAppDetectors(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().listSiteDetectors("Sample-WestUSResourceGroup", "SampleApp",
            "availability", com.azure.core.util.Context.NONE);
    }
}
