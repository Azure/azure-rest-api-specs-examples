import com.azure.resourcemanager.trafficmanager.models.EndpointTypes;

/** Samples for Endpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Endpoint-GET-External-WithLocation.json
     */
    /**
     * Sample code: Endpoint-GET-External-WithLocation.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointGETExternalWithLocation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .trafficManagerProfiles()
            .manager()
            .serviceClient()
            .getEndpoints()
            .getWithResponse(
                "azuresdkfornetautoresttrafficmanager1421",
                "azsmnet6386",
                EndpointTypes.EXTERNAL_ENDPOINTS,
                "azsmnet7187",
                com.azure.core.util.Context.NONE);
    }
}
