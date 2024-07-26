
/**
 * Samples for AppServiceEnvironments ListDiagnostics.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/
     * AppServiceEnvironments_ListDiagnostics.json
     */
    /**
     * Sample code: Get diagnostic information for an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getDiagnosticInformationForAnAppServiceEnvironment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().listDiagnosticsWithResponse("test-rg",
            "test-ase", com.azure.core.util.Context.NONE);
    }
}
