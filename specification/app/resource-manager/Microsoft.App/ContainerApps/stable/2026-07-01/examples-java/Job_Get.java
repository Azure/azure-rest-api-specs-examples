
/**
 * Samples for Jobs GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Job_Get.json
     */
    /**
     * Sample code: Get Container Apps Job.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getContainerAppsJob(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().getByResourceGroupWithResponse("rg", "testcontainerAppsJob0", com.azure.core.util.Context.NONE);
    }
}
