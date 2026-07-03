
/**
 * Samples for AvailableEndpointServices List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/EndpointServicesList.json
     */
    /**
     * Sample code: EndpointServicesList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void endpointServicesList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAvailableEndpointServices().list("westus", com.azure.core.util.Context.NONE);
    }
}
