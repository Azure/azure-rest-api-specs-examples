
/**
 * Samples for Diagnostics GetSiteDetectorResponse.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/Diagnostics_GetSiteDetectorResponse.
     * json
     */
    /**
     * Sample code: Get App Detector Response.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppDetectorResponse(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().getSiteDetectorResponseWithResponse(
            "Sample-WestUSResourceGroup", "SampleApp", "runtimeavailability", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
