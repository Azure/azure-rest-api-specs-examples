/** Samples for ContainerAppsAuthConfigs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/AuthConfigs_Delete.json
     */
    /**
     * Sample code: Delete Container App AuthConfig.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteContainerAppAuthConfig(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsAuthConfigs()
            .deleteWithResponse("workerapps-rg-xj", "testcanadacentral", "current", com.azure.core.util.Context.NONE);
    }
}
