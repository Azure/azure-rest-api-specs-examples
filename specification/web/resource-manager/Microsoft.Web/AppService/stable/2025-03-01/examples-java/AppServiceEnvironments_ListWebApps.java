
/**
 * Samples for AppServiceEnvironments ListWebApps.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_ListWebApps.json
     */
    /**
     * Sample code: Get all apps in an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAllAppsInAnAppServiceEnvironment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().listWebApps("test-rg", "test-ase", null,
            com.azure.core.util.Context.NONE);
    }
}
