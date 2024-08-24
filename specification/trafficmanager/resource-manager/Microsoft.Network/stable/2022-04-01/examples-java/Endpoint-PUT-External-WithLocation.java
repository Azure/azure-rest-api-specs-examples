
import com.azure.resourcemanager.trafficmanager.fluent.models.EndpointInner;
import com.azure.resourcemanager.trafficmanager.models.EndpointStatus;
import com.azure.resourcemanager.trafficmanager.models.EndpointTypes;

/**
 * Samples for Endpoints CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Endpoint-PUT-External-
     * WithLocation.json
     */
    /**
     * Sample code: Endpoint-PUT-External-WithLocation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointPUTExternalWithLocation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.trafficManagerProfiles().manager().serviceClient().getEndpoints().createOrUpdateWithResponse(
            "azuresdkfornetautoresttrafficmanager1421", "azsmnet6386", EndpointTypes.EXTERNAL_ENDPOINTS, "azsmnet7187",
            new EndpointInner().withName("azsmnet7187")
                .withType("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints").withTarget("foobar.contoso.com")
                .withEndpointStatus(EndpointStatus.ENABLED).withEndpointLocation("North Europe"),
            com.azure.core.util.Context.NONE);
    }
}
