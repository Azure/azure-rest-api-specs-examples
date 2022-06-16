import com.azure.core.util.Context;

/** Samples for ContainerAppsAuthConfigs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/AuthConfigs_Delete.json
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
            .deleteWithResponse("workerapps-rg-xj", "testcanadacentral", "current", Context.NONE);
    }
}
