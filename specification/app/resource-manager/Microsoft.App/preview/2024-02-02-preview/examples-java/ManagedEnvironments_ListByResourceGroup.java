
/**
 * Samples for ManagedEnvironments ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * ManagedEnvironments_ListByResourceGroup.json
     */
    /**
     * Sample code: List environments by resource group.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listEnvironmentsByResourceGroup(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironments().listByResourceGroup("examplerg", com.azure.core.util.Context.NONE);
    }
}
