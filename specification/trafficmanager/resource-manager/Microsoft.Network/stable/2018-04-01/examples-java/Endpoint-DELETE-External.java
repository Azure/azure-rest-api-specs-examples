import com.azure.core.util.Context;

/** Samples for Endpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-04-01/examples/Endpoint-DELETE-External.json
     */
    /**
     * Sample code: Endpoint-DELETE-External.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointDELETEExternal(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .trafficManagerProfiles()
            .manager()
            .serviceClient()
            .getEndpoints()
            .deleteWithResponse(
                "azuresdkfornetautoresttrafficmanager1421",
                "azsmnet6386",
                "ExternalEndpoints",
                "azsmnet7187",
                Context.NONE);
    }
}
