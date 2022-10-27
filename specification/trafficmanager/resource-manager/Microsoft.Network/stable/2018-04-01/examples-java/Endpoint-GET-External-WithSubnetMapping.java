import com.azure.core.util.Context;

/** Samples for Endpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-04-01/examples/Endpoint-GET-External-WithSubnetMapping.json
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
                "ExternalEndpoints",
                "My%20external%20endpoint",
                Context.NONE);
    }
}
