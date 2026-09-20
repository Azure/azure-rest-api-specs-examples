
/**
 * Samples for ContainerAppPrivateLinkResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerAppPrivateLinkResources_List.json
     */
    /**
     * Sample code: List Private Link Resources by Container App.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listPrivateLinkResourcesByContainerApp(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppPrivateLinkResources().list("examplerg", "testcontainerapp0",
            com.azure.core.util.Context.NONE);
    }
}
