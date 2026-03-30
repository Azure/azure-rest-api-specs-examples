
/**
 * Samples for Diagnostics ListHostingEnvironmentDetectorResponses.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ListHostingEnvironmentDetectorResponses.json
     */
    /**
     * Sample code: Get App Service Environment Detector Responses.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAppServiceEnvironmentDetectorResponses(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().listHostingEnvironmentDetectorResponses("Sample-WestUSResourceGroup",
            "SampleAppServiceEnvironment", com.azure.core.util.Context.NONE);
    }
}
