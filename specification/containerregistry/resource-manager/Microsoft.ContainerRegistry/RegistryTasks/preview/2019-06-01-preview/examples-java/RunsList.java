
/**
 * Samples for Runs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/RunsList.json
     */
    /**
     * Sample code: Runs_List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void runsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRuns().list("myResourceGroup", "myRegistry", "", 10,
            com.azure.core.util.Context.NONE);
    }
}
