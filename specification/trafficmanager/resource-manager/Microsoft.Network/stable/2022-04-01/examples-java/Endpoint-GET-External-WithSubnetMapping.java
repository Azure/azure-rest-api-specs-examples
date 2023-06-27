import com.azure.resourcemanager.trafficmanager.models.EndpointTypes;

/** Samples for Endpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Endpoint-GET-External-WithSubnetMapping.json
     */
    /**
     * Sample code: Endpoint-GET-External-WithSubnetMapping.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointGETExternalWithSubnetMapping(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .trafficManagerProfiles()
            .manager()
            .serviceClient()
            .getEndpoints()
            .getWithResponse(
                "azuresdkfornetautoresttrafficmanager2191",
                "azuresdkfornetautoresttrafficmanager8224",
                EndpointTypes.EXTERNAL_ENDPOINTS,
                "My%20external%20endpoint",
                com.azure.core.util.Context.NONE);
    }
}
