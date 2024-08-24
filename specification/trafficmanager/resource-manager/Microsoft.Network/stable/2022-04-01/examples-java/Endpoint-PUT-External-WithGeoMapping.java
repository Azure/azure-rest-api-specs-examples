
import com.azure.resourcemanager.trafficmanager.fluent.models.EndpointInner;
import com.azure.resourcemanager.trafficmanager.models.EndpointStatus;
import com.azure.resourcemanager.trafficmanager.models.EndpointTypes;
import java.util.Arrays;

/**
 * Samples for Endpoints CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Endpoint-PUT-External-
     * WithGeoMapping.json
     */
    /**
     * Sample code: Endpoint-PUT-External-WithGeoMapping.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointPUTExternalWithGeoMapping(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.trafficManagerProfiles().manager().serviceClient().getEndpoints().createOrUpdateWithResponse(
            "azuresdkfornetautoresttrafficmanager2191", "azuresdkfornetautoresttrafficmanager8224",
            EndpointTypes.EXTERNAL_ENDPOINTS, "My%20external%20endpoint",
            new EndpointInner().withName("My external endpoint")
                .withType("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints").withTarget("foobar.contoso.com")
                .withEndpointStatus(EndpointStatus.ENABLED).withGeoMapping(Arrays.asList("GEO-AS", "GEO-AF")),
            com.azure.core.util.Context.NONE);
    }
}
