
/**
 * Samples for ManagedEnvironmentDiagnostics GetDetector.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironmentDiagnostics_Get.json
     */
    /**
     * Sample code: Get diagnostic data for a managed environments.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getDiagnosticDataForAManagedEnvironments(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironmentDiagnostics().getDetectorWithResponse("mikono-workerapp-test-rg", "mikonokubeenv",
            "ManagedEnvAvailabilityMetrics", com.azure.core.util.Context.NONE);
    }
}
