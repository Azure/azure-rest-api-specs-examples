
/**
 * Samples for ManagedEnvironments ListWorkloadProfileStates.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * ManagedEnvironments_ListWorkloadProfileStates.json
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
