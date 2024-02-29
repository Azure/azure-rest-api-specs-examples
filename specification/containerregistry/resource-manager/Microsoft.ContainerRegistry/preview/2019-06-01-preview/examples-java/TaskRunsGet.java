
/**
 * Samples for TaskRuns Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/
     * TaskRunsGet.json
     */
    /**
     * Sample code: TaskRuns_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void taskRunsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTaskRuns().getWithResponse("myResourceGroup",
            "myRegistry", "myRun", com.azure.core.util.Context.NONE);
    }
}
