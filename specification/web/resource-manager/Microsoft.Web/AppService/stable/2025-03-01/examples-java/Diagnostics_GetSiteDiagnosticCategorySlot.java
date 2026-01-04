
/**
 * Samples for Diagnostics GetSiteDiagnosticCategory.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * Diagnostics_GetSiteDiagnosticCategorySlot.json
     */
    /**
     * Sample code: Get App Slot Diagnostic Category.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppSlotDiagnosticCategory(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().getSiteDiagnosticCategoryWithResponse(
            "Sample-WestUSResourceGroup", "SampleApp", "availability", com.azure.core.util.Context.NONE);
    }
}
