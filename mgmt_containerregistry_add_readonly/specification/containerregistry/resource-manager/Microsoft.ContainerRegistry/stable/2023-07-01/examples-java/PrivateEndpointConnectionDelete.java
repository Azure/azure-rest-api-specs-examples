
/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * mgmt_containerregistry_add_readonly/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/
     * stable/2023-07-01/examples/PrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: PrivateEndpointConnectionDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void privateEndpointConnectionDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getPrivateEndpointConnections().delete("myResourceGroup",
            "myRegistry", "myConnection", com.azure.core.util.Context.NONE);
    }
}
