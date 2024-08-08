
/**
 * Samples for Jobs GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/Job_Get.json
     */
    /**
     * Sample code: Get Container Apps Job.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getContainerAppsJob(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().getByResourceGroupWithResponse("rg", "testcontainerappsjob0", com.azure.core.util.Context.NONE);
    }
}
