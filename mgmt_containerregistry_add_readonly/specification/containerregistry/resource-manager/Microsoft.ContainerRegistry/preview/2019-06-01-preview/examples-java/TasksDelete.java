
/** Samples for Tasks Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * mgmt_containerregistry_add_readonly/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/
     * preview/2019-06-01-preview/examples/TasksDelete.json
     */
    /**
     * Sample code: Tasks_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tasksDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTasks().delete("myResourceGroup", "myRegistry",
            "myTask", com.azure.core.util.Context.NONE);
    }
}
