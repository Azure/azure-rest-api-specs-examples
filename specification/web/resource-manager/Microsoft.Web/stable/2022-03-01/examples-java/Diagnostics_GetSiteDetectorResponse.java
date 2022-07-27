import com.azure.core.util.Context;

/** Samples for Diagnostics GetSiteDetectorResponseSlot. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/Diagnostics_GetSiteDetectorResponse.json
     */
    /**
     * Sample code: Get App Detector Response.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppDetectorResponse(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .getSiteDetectorResponseSlotWithResponse(
                "Sample-WestUSResourceGroup",
                "SampleApp",
                "runtimeavailability",
                "staging",
                null,
                null,
                null,
                Context.NONE);
    }
}
