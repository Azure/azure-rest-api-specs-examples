
/**
 * Samples for AppServiceEnvironments ListDiagnostics.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListDiagnostics.json
     */
    /**
     * Sample code: Get diagnostic information for an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDiagnosticInformationForAnAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listDiagnosticsWithResponse("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
