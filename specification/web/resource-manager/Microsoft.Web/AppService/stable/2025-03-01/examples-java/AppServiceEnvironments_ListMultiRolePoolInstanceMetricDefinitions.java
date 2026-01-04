
/**
 * Samples for AppServiceEnvironments ListMultiRolePoolInstanceMetricDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_ListMultiRolePoolInstanceMetricDefinitions.json
     */
    /**
     * Sample code: Get metric definitions for a specific instance of a multi-role pool of an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getMetricDefinitionsForASpecificInstanceOfAMultiRolePoolOfAnAppServiceEnvironment(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments()
            .listMultiRolePoolInstanceMetricDefinitions("test-rg", "test-ase", "10.7.1.8",
                com.azure.core.util.Context.NONE);
    }
}
