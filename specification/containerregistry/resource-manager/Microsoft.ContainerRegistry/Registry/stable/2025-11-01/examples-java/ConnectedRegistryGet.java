
/**
 * Samples for ConnectedRegistries Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * ConnectedRegistryGet.json
     */
    /**
     * Sample code: ConnectedRegistryGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void connectedRegistryGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getConnectedRegistries()
            .getWithResponse("myResourceGroup", "myRegistry", "myConnectedRegistry", com.azure.core.util.Context.NONE);
    }
}
