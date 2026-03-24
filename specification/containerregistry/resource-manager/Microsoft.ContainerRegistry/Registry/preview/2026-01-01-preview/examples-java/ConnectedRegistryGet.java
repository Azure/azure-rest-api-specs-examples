
/**
 * Samples for ConnectedRegistries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ConnectedRegistryGet.json
     */
    /**
     * Sample code: ConnectedRegistryGet.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        connectedRegistryGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getConnectedRegistries().getWithResponse("myResourceGroup", "myRegistry",
            "myConnectedRegistry", com.azure.core.util.Context.NONE);
    }
}
