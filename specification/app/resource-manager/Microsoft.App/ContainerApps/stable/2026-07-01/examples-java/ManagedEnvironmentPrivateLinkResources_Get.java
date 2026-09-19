
/**
 * Samples for ManagedEnvironmentPrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironmentPrivateLinkResources_Get.json
     */
    /**
     * Sample code: Get a Private Link Resource by Managed Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getAPrivateLinkResourceByManagedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironmentPrivateLinkResources().getWithResponse("examplerg", "managedEnv",
            "managedEnvironments", com.azure.core.util.Context.NONE);
    }
}
