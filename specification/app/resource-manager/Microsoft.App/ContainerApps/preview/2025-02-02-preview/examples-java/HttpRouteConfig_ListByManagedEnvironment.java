
/**
 * Samples for HttpRouteConfig List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * HttpRouteConfig_ListByManagedEnvironment.json
     */
    /**
     * Sample code: List Managed Http Routes by Managed Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listManagedHttpRoutesByManagedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.httpRouteConfigs().list("examplerg", "testcontainerenv", com.azure.core.util.Context.NONE);
    }
}
