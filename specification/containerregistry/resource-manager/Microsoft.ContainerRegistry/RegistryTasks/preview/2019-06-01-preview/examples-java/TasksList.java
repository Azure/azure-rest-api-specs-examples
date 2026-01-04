
/**
 * Samples for Tasks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/TasksList.json
     */
    /**
     * Sample code: Tasks_List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tasksList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTasks().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
