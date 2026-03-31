
/**
 * Samples for AppServiceEnvironments ListOperations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListOperations.json
     */
    /**
     * Sample code: List all currently running operations on the App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAllCurrentlyRunningOperationsOnTheAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listOperationsWithResponse("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
