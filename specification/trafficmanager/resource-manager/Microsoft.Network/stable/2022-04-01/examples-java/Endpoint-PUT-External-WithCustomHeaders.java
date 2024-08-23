
import com.azure.resourcemanager.trafficmanager.fluent.models.EndpointInner;
import com.azure.resourcemanager.trafficmanager.models.EndpointPropertiesCustomHeadersItem;
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
     * WithCustomHeaders.json
     */
    /**
     * Sample code: Endpoint-PUT-External-WithCustomHeaders.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointPUTExternalWithCustomHeaders(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.trafficManagerProfiles().manager().serviceClient().getEndpoints().createOrUpdateWithResponse(
            "azuresdkfornetautoresttrafficmanager1421", "azsmnet6386", EndpointTypes.EXTERNAL_ENDPOINTS, "azsmnet7187",
            new EndpointInner().withName("azsmnet7187")
                .withType("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints").withTarget("foobar.contoso.com")
                .withEndpointStatus(EndpointStatus.ENABLED).withEndpointLocation("North Europe").withCustomHeaders(
                    Arrays.asList(new EndpointPropertiesCustomHeadersItem().withName("header-1").withValue("value-1"),
                        new EndpointPropertiesCustomHeadersItem().withName("header-2").withValue("value-2"))),
            com.azure.core.util.Context.NONE);
    }
}
