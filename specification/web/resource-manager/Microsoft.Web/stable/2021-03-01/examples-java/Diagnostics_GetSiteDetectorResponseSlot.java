import com.azure.core.util.Context;

/** Samples for Diagnostics GetSiteDetectorResponse. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/Diagnostics_GetSiteDetectorResponseSlot.json
     */
    /**
     * Sample code: Get App Slot Detector Response.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppSlotDetectorResponse(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .getSiteDetectorResponseWithResponse(
                "Sample-WestUSResourceGroup", "SampleApp", "runtimeavailability", null, null, null, Context.NONE);
    }
}
