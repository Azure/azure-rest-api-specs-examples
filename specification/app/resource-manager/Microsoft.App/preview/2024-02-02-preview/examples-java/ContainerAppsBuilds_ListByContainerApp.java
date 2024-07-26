
/**
 * Samples for ContainerAppsBuildsByContainerApp List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * ContainerAppsBuilds_ListByContainerApp.json
     */
    /**
     * Sample code: ContainerAppsBuilds_ListByContainerApp_0.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void containerAppsBuildsListByContainerApp0(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsBuildsByContainerApps().list("rg", "testCapp", com.azure.core.util.Context.NONE);
    }
}
