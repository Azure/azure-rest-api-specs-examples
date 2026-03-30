
/**
 * Samples for Diagnostics GetSiteDiagnosticCategory.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetSiteDiagnosticCategory.json
     */
    /**
     * Sample code: Get App Diagnostic Category.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppDiagnosticCategory(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getSiteDiagnosticCategoryWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "availability", com.azure.core.util.Context.NONE);
    }
}
