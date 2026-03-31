
/**
 * Samples for Diagnostics ListSiteDetectorResponses.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ListSiteDetectorResponses.json
     */
    /**
     * Sample code: Get App Detector Responses.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppDetectorResponses(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().listSiteDetectorResponses("Sample-WestUSResourceGroup", "SampleApp",
            com.azure.core.util.Context.NONE);
    }
}
