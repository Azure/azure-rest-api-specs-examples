
/**
 * Samples for Jobs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/Job_Delete.json
     */
    /**
     * Sample code: Delete Container Apps Job.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteContainerAppsJob(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().delete("rg", "testworkercontainerappsjob0", com.azure.core.util.Context.NONE);
    }
}
