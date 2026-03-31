
/**
 * Samples for AppServiceEnvironments ListWorkerPools.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListWorkerPools.json
     */
    /**
     * Sample code: Get all worker pools of an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAllWorkerPoolsOfAnAppServiceEnvironment(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listWorkerPools("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
