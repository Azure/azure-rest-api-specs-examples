import com.azure.core.util.Context;

/** Samples for Diagnostics GetSiteDetector. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/Diagnostics_GetSiteDetectorSlot.json
     */
    /**
     * Sample code: Get App Slot Detector.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppSlotDetector(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .getSiteDetectorWithResponse(
                "Sample-WestUSResourceGroup", "SampleApp", "availability", "sitecrashes", Context.NONE);
    }
}
