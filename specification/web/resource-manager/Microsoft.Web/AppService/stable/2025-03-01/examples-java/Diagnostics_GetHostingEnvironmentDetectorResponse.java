
/**
 * Samples for Diagnostics GetHostingEnvironmentDetectorResponse.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * Diagnostics_GetHostingEnvironmentDetectorResponse.json
     */
    /**
     * Sample code: Get App Service Environment Detector Responses.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppServiceEnvironmentDetectorResponses(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().getHostingEnvironmentDetectorResponseWithResponse(
            "Sample-WestUSResourceGroup", "SampleAppServiceEnvironment", "runtimeavailability", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
