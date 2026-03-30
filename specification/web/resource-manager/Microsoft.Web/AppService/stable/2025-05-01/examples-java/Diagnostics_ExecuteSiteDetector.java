
/**
 * Samples for Diagnostics ExecuteSiteDetector.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ExecuteSiteDetector.json
     */
    /**
     * Sample code: Execute site detector.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void executeSiteDetector(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().executeSiteDetectorWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "sitecrashes", "availability", null, null, null, com.azure.core.util.Context.NONE);
    }
}
