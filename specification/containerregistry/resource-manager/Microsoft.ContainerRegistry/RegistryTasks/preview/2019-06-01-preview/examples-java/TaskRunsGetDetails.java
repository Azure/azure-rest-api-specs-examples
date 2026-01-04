
/**
 * Samples for TaskRuns GetDetails.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/TaskRunsGetDetails.json
     */
    /**
     * Sample code: TaskRuns_GetDetails.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void taskRunsGetDetails(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTaskRuns().getDetailsWithResponse("myResourceGroup",
            "myRegistry", "myRun", com.azure.core.util.Context.NONE);
    }
}
