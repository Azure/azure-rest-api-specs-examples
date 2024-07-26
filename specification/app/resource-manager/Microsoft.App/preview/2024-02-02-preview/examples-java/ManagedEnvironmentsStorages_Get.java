
/**
 * Samples for ManagedEnvironmentsStorages Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
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
