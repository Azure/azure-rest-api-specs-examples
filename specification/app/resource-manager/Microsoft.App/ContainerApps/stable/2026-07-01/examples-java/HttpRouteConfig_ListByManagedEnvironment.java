
/**
 * Samples for HttpRouteConfig List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/HttpRouteConfig_ListByManagedEnvironment.json
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
