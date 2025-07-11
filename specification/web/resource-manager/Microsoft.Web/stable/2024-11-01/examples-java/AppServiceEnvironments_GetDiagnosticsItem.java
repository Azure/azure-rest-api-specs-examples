
/**
 * Samples for AppServiceEnvironments GetDiagnosticsItem.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/
     * AppServiceEnvironments_GetDiagnosticsItem.json
     */
    /**
     * Sample code: Get a diagnostics item for an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getADiagnosticsItemForAnAppServiceEnvironment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().getDiagnosticsItemWithResponse("test-rg",
            "test-ase", "test-diagnostic", com.azure.core.util.Context.NONE);
    }
}
