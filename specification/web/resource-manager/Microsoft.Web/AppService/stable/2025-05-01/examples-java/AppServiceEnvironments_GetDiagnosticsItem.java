
/**
 * Samples for AppServiceEnvironments GetDiagnosticsItem.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_GetDiagnosticsItem.json
     */
    /**
     * Sample code: Get a diagnostics item for an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getADiagnosticsItemForAnAppServiceEnvironment(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().getDiagnosticsItemWithResponse("test-rg", "test-ase",
            "test-diagnostic", com.azure.core.util.Context.NONE);
    }
}
