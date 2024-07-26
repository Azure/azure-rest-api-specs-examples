
/**
 * Samples for ContainerAppsPatches ListByContainerApp.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * ContainerAppsPatches_ListByContainerApp.json
     */
    /**
     * Sample code: ContainerAppsPatches_ListByContainerApp_0.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void containerAppsPatchesListByContainerApp0(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsPatches().listByContainerApp("rg", "test-app", null, com.azure.core.util.Context.NONE);
    }
}
