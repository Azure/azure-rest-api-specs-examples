
import com.azure.resourcemanager.trafficmanager.fluent.models.EndpointInner;
import com.azure.resourcemanager.trafficmanager.models.EndpointPropertiesSubnetsItem;
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
     * WithSubnetMapping.json
     */
    /**
     * Sample code: Endpoint-PUT-External-WithSubnetMapping.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointPUTExternalWithSubnetMapping(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.trafficManagerProfiles().manager().serviceClient().getEndpoints().createOrUpdateWithResponse(
            "azuresdkfornetautoresttrafficmanager2191", "azuresdkfornetautoresttrafficmanager8224",
            EndpointTypes.EXTERNAL_ENDPOINTS, "My%20external%20endpoint",
            new EndpointInner().withName("My external endpoint")
                .withType("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints").withTarget("foobar.contoso.com")
                .withEndpointStatus(EndpointStatus.ENABLED)
                .withSubnets(Arrays.asList(new EndpointPropertiesSubnetsItem().withFirst("1.2.3.0").withScope(24),
                    new EndpointPropertiesSubnetsItem().withFirst("25.26.27.28").withLast("29.30.31.32"))),
            com.azure.core.util.Context.NONE);
    }
}
