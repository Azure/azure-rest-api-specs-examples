import com.azure.core.util.Context;

/** Samples for AvailableEndpointServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/EndpointServicesList.json
     */
    /**
     * Sample code: EndpointServicesList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointServicesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAvailableEndpointServices().list("westus", Context.NONE);
    }
}
