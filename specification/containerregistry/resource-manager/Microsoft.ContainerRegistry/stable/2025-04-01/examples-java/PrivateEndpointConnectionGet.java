
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2025-04-01/examples/
     * PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: PrivateEndpointConnectionGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void privateEndpointConnectionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getPrivateEndpointConnections()
            .getWithResponse("myResourceGroup", "myRegistry", "myConnection", com.azure.core.util.Context.NONE);
    }
}
