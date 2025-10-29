
/**
 * Samples for HttpRouteConfig List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/
     * HttpRouteConfig_ListByManagedEnvironment.json
     */
    /**
     * Sample code: List Http Route Configs by Managed Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listHttpRouteConfigsByManagedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.httpRouteConfigs().list("examplerg", "testcontainerenv", com.azure.core.util.Context.NONE);
    }
}
