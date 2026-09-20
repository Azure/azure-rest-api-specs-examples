
/**
 * Samples for ConnectedEnvironments ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ConnectedEnvironments_ListByResourceGroup.json
     */
    /**
     * Sample code: List environments by resource group.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listEnvironmentsByResourceGroup(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironments().listByResourceGroup("examplerg", com.azure.core.util.Context.NONE);
    }
}
