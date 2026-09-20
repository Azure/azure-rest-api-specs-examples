
/**
 * Samples for ManagedEnvironmentUsages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironmentUsages_List.json
     */
    /**
     * Sample code: List managed environment usages.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listManagedEnvironmentUsages(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironmentUsages().list("examplerg", "jlaw-demo1", com.azure.core.util.Context.NONE);
    }
}
