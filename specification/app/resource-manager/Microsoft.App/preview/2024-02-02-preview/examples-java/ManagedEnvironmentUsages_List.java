
/**
 * Samples for ManagedEnvironmentUsages List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * ManagedEnvironmentUsages_List.json
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
