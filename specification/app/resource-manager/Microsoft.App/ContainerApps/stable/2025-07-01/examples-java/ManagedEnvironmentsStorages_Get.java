
/**
 * Samples for ManagedEnvironmentsStorages Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/
     * ManagedEnvironmentsStorages_Get.json
     */
    /**
     * Sample code: get a environments storage.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getAEnvironmentsStorage(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironmentsStorages().getWithResponse("examplerg", "managedEnv", "jlaw-demo1",
            com.azure.core.util.Context.NONE);
    }
}
