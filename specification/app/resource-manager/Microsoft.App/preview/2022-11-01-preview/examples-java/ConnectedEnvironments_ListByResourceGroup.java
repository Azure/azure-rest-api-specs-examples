/** Samples for ConnectedEnvironments ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ConnectedEnvironments_ListByResourceGroup.json
     */
    /**
     * Sample code: List environments by resource group.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listEnvironmentsByResourceGroup(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironments().listByResourceGroup("examplerg", com.azure.core.util.Context.NONE);
    }
}
