
import com.azure.resourcemanager.trafficmanager.fluent.models.EndpointInner;
import com.azure.resourcemanager.trafficmanager.models.AlwaysServe;
import com.azure.resourcemanager.trafficmanager.models.EndpointStatus;
import com.azure.resourcemanager.trafficmanager.models.EndpointTypes;

/**
 * Samples for Endpoints CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/Endpoint-PUT-External-WithAlwaysServe.json
     */
    /**
     * Sample code: Endpoint-PUT-External-WithAlwaysServe.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void
        endpointPUTExternalWithAlwaysServe(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getEndpoints().createOrUpdateWithResponse("azuresdkfornetautoresttrafficmanager1421",
            "azsmnet6386", EndpointTypes.EXTERNAL_ENDPOINTS, "azsmnet7187",
            new EndpointInner().withName("azsmnet7187")
                .withType("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints").withTarget("foobar.contoso.com")
                .withEndpointStatus(EndpointStatus.ENABLED).withEndpointLocation("North Europe")
                .withAlwaysServe(AlwaysServe.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
