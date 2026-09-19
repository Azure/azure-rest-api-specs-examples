
/**
 * Samples for ContainerAppsAuthConfigs ListByContainerApp.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/AuthConfigs_ListByContainer.json
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
