
/**
 * Samples for AppServiceEnvironments ListCapacities.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListCapacities.json
     */
    /**
     * Sample code: Get the used, available, and total worker capacity an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getTheUsedAvailableAndTotalWorkerCapacityAnAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listCapacities("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
