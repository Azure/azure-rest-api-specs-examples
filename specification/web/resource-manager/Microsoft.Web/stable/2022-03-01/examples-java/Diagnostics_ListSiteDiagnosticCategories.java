import com.azure.core.util.Context;

/** Samples for Diagnostics ListSiteDiagnosticCategories. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/Diagnostics_ListSiteDiagnosticCategories.json
     */
    /**
     * Sample code: List App Diagnostic Categories.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppDiagnosticCategories(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .listSiteDiagnosticCategories("Sample-WestUSResourceGroup", "SampleApp", Context.NONE);
    }
}
