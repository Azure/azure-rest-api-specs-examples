import com.azure.core.util.Context;

/** Samples for ContainerAppsAuthConfigs ListByContainerApp. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/AuthConfigs_ListByContainer.json
     */
    /**
     * Sample code: List Auth Configs by Container Apps.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listAuthConfigsByContainerApps(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsAuthConfigs().listByContainerApp("workerapps-rg-xj", "testcanadacentral", Context.NONE);
    }
}
