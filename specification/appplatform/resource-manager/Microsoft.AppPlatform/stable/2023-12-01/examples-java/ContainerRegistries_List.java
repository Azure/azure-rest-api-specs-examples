
/**
 * Samples for ContainerRegistries List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * ContainerRegistries_List.json
     */
    /**
     * Sample code: ContainerRegistries_List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerRegistriesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getContainerRegistries().list("myResourceGroup", "my-service",
            com.azure.core.util.Context.NONE);
    }
}
