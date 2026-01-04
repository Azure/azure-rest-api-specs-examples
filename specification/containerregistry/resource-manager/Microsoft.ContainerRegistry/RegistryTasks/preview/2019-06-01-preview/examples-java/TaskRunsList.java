
/**
 * Samples for TaskRuns List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/TaskRunsList.json
     */
    /**
     * Sample code: TaskRuns_List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void taskRunsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTaskRuns().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
