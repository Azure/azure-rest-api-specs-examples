import com.azure.core.util.Context;

/** Samples for PrivateEndpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/PrivateEndpointDelete.json
     */
    /**
     * Sample code: Delete private endpoint.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletePrivateEndpoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPrivateEndpoints().delete("rg1", "testPe", Context.NONE);
    }
}
