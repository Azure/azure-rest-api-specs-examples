
/**
 * Samples for ContainerAppPrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerAppPrivateLinkResources_Get.json
     */
    /**
     * Sample code: Get a Private Link Resource by Container App.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getAPrivateLinkResourceByContainerApp(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppPrivateLinkResources().getWithResponse("examplerg", "testcontainerapp0", "containerApps",
            com.azure.core.util.Context.NONE);
    }
}
