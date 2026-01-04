
/**
 * Samples for TaskRuns Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/TaskRunsDelete.json
     */
    /**
     * Sample code: TaskRuns_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void taskRunsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTaskRuns().delete("myResourceGroup", "myRegistry",
            "myRun", com.azure.core.util.Context.NONE);
    }
}
