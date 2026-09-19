
/**
 * Samples for ManagedEnvironmentDiagnostics ListDetectors.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironmentDiagnostics_List.json
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
