
import com.azure.resourcemanager.trafficmanager.models.EndpointTypes;

/** Samples for Endpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Endpoint-DELETE-
     * External.json
     */
    /**
     * Sample code: Endpoint-DELETE-External.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointDELETEExternal(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.trafficManagerProfiles().manager().serviceClient().getEndpoints().deleteWithResponse(
            "azuresdkfornetautoresttrafficmanager1421", "azsmnet6386", EndpointTypes.EXTERNAL_ENDPOINTS, "azsmnet7187",
            com.azure.core.util.Context.NONE);
    }
}
