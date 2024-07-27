
/**
 * Samples for AppServiceEnvironments ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/AppServiceEnvironments_ListUsages.
     * json
     */
    /**
     * Sample code: Get global usage metrics of an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getGlobalUsageMetricsOfAnAppServiceEnvironment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().listUsages("test-rg", "test-ase", null,
            com.azure.core.util.Context.NONE);
    }
}
