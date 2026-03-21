
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: PrivateEndpointConnectionGet.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        privateEndpointConnectionGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().getWithResponse("myResourceGroup", "myRegistry",
            "myConnection", com.azure.core.util.Context.NONE);
    }
}
