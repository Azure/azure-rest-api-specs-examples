
/**
 * Samples for ManagedEnvironmentDiagnostics GetDetector.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/
     * ManagedEnvironmentDiagnostics_Get.json
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
