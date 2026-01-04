
/**
 * Samples for Diagnostics GetSiteDiagnosticCategory.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * Diagnostics_GetSiteDiagnosticCategory.json
     */
    /**
     * Sample code: Get App Diagnostic Category.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppDiagnosticCategory(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().getSiteDiagnosticCategoryWithResponse(
            "Sample-WestUSResourceGroup", "SampleApp", "availability", com.azure.core.util.Context.NONE);
    }
}
