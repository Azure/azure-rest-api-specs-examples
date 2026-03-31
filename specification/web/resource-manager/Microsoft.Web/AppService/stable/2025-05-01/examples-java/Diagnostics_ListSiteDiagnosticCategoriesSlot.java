
/**
 * Samples for Diagnostics ListSiteDiagnosticCategories.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ListSiteDiagnosticCategoriesSlot.json
     */
    /**
     * Sample code: List App Slot Diagnostic Categories.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAppSlotDiagnosticCategories(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().listSiteDiagnosticCategories("Sample-WestUSResourceGroup", "SampleApp",
            com.azure.core.util.Context.NONE);
    }
}
