
/**
 * Samples for Diagnostics ListSiteDiagnosticCategoriesSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * Diagnostics_ListSiteDiagnosticCategories.json
     */
    /**
     * Sample code: List App Diagnostic Categories.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppDiagnosticCategories(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().listSiteDiagnosticCategoriesSlot(
            "Sample-WestUSResourceGroup", "SampleApp", "Production", com.azure.core.util.Context.NONE);
    }
}
