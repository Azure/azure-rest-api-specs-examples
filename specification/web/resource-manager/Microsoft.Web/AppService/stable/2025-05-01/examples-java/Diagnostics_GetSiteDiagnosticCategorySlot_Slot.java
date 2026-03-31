
/**
 * Samples for Diagnostics GetSiteDiagnosticCategorySlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetSiteDiagnosticCategorySlot_Slot.json
     */
    /**
     * Sample code: Get App Slot Diagnostic Category.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppSlotDiagnosticCategory(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getSiteDiagnosticCategorySlotWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "availability", "staging", com.azure.core.util.Context.NONE);
    }
}
