
/**
 * Samples for ManagedEnvironmentDiagnostics ListDetectors.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ManagedEnvironmentDiagnostics_List.
     * json
     */
    /**
     * Sample code: Get the list of available diagnostic data for a managed environments.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getTheListOfAvailableDiagnosticDataForAManagedEnvironments(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironmentDiagnostics().listDetectorsWithResponse("mikono-workerapp-test-rg", "mikonokubeenv",
            com.azure.core.util.Context.NONE);
    }
}
