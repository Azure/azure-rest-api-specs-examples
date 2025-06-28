
/**
 * Samples for AppServiceEnvironments ListWorkerPools.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/
     * AppServiceEnvironments_ListWorkerPools.json
     */
    /**
     * Sample code: Get all worker pools of an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAllWorkerPoolsOfAnAppServiceEnvironment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().listWorkerPools("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
