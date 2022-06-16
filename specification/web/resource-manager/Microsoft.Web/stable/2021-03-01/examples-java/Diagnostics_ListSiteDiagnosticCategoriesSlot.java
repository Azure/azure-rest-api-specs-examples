import com.azure.core.util.Context;

/** Samples for Diagnostics ListSiteDiagnosticCategoriesSlot. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/Diagnostics_ListSiteDiagnosticCategoriesSlot.json
     */
    /**
     * Sample code: List App Slot Diagnostic Categories.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppSlotDiagnosticCategories(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .listSiteDiagnosticCategoriesSlot("Sample-WestUSResourceGroup", "SampleApp", "staging", Context.NONE);
    }
}
