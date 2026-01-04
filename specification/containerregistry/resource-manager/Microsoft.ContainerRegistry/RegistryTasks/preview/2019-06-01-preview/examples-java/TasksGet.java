
/**
 * Samples for Tasks Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/TasksGet.json
     */
    /**
     * Sample code: Tasks_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tasksGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTasks().getWithResponse("myResourceGroup",
            "myRegistry", "myTask", com.azure.core.util.Context.NONE);
    }
}
