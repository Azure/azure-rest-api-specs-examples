import com.azure.core.util.Context;

/** Samples for Tasks GetDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TasksGetDetails.json
     */
    /**
     * Sample code: Tasks_GetDetails.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tasksGetDetails(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getTasks()
            .getDetailsWithResponse("myResourceGroup", "myRegistry", "myTask", Context.NONE);
    }
}
