
/**
 * Samples for AppServiceEnvironments ListAppServicePlans.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_ListAppServicePlans.json
     */
    /**
     * Sample code: Get all App Service plans in an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAllAppServicePlansInAnAppServiceEnvironment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().listAppServicePlans("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
