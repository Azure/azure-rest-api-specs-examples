
/**
 * Samples for Diagnostics GetHostingEnvironmentDetectorResponse.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetHostingEnvironmentDetectorResponse.json
     */
    /**
     * Sample code: Get App Service Environment Detector Responses.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAppServiceEnvironmentDetectorResponses(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getHostingEnvironmentDetectorResponseWithResponse(
            "Sample-WestUSResourceGroup", "SampleAppServiceEnvironment", "runtimeavailability", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
