/** Samples for Diagnostics GetSiteDiagnosticCategorySlot. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/Diagnostics_GetSiteDiagnosticCategorySlot.json
     */
    /**
     * Sample code: Get App Slot Diagnostic Category.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppSlotDiagnosticCategory(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .getSiteDiagnosticCategorySlotWithResponse(
                "Sample-WestUSResourceGroup", "SampleApp", "availability", "staging", com.azure.core.util.Context.NONE);
    }
}
