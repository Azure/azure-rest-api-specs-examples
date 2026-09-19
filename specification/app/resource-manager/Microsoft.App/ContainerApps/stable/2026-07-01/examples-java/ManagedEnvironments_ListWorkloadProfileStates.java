
/**
 * Samples for ManagedEnvironments ListWorkloadProfileStates.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironments_ListWorkloadProfileStates.json
     */
    /**
     * Sample code: List environments by subscription.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listEnvironmentsBySubscription(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironments().listWorkloadProfileStates("examplerg", "jlaw-demo1",
            com.azure.core.util.Context.NONE);
    }
}
