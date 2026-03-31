
/**
 * Samples for AppServiceEnvironments ListWebApps.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListWebApps.json
     */
    /**
     * Sample code: Get all apps in an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAllAppsInAnAppServiceEnvironment(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listWebApps("test-rg", "test-ase", null,
            com.azure.core.util.Context.NONE);
    }
}
