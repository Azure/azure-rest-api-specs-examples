
/**
 * Samples for Diagnostics GetSiteDetector.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetSiteDetector.json
     */
    /**
     * Sample code: Get App Detector.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppDetector(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getSiteDetectorWithResponse("Sample-WestUSResourceGroup", "SampleApp",
            "availability", "sitecrashes", com.azure.core.util.Context.NONE);
    }
}
