import com.azure.core.util.Context;

/** Samples for Endpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-04-01/examples/Endpoint-GET-External-WithLocation.json
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
                "ExternalEndpoints",
                "azsmnet7187",
                Context.NONE);
    }
}
