
/**
 * Samples for HttpRouteConfig Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/HttpRouteConfig_Delete.
     * json
     */
    /**
     * Sample code: Delete a Http Route Config.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteAHttpRouteConfig(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.httpRouteConfigs().delete("examplerg", "testcontainerenv", "httproutefriendlyname",
            com.azure.core.util.Context.NONE);
    }
}
