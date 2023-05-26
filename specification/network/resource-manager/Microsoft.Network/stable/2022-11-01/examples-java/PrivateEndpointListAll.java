/** Samples for PrivateEndpoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/PrivateEndpointListAll.json
     */
    /**
     * Sample code: List all private endpoints.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllPrivateEndpoints(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPrivateEndpoints().list(com.azure.core.util.Context.NONE);
    }
}
