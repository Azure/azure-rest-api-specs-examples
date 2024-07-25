
/**
 * Samples for ContainerAppsAuthConfigs ListByContainerApp.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/AuthConfigs_ListByContainer.
     * json
     */
    /**
     * Sample code: List Auth Configs by Container Apps.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listAuthConfigsByContainerApps(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsAuthConfigs().listByContainerApp("workerapps-rg-xj", "testcanadacentral",
            com.azure.core.util.Context.NONE);
    }
}
