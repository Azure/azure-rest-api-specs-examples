
/**
 * Samples for Diagnostics GetSiteDiagnosticCategory.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetSiteDiagnosticCategorySlot.json
     */
    /**
     * Sample code: Get App Slot Diagnostic Category.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppSlotDiagnosticCategory(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getSiteDiagnosticCategoryWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "availability", com.azure.core.util.Context.NONE);
    }
}
