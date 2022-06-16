import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: PrivateEndpointConnectionGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void privateEndpointConnectionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .getWithResponse("myResourceGroup", "myRegistry", "myConnection", Context.NONE);
    }
}
