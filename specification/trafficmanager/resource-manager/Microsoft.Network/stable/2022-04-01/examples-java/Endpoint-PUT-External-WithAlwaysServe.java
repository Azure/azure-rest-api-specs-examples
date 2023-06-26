import com.azure.resourcemanager.trafficmanager.fluent.models.EndpointInner;
import com.azure.resourcemanager.trafficmanager.models.AlwaysServe;
import com.azure.resourcemanager.trafficmanager.models.EndpointStatus;
import com.azure.resourcemanager.trafficmanager.models.EndpointTypes;

/** Samples for Endpoints CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Endpoint-PUT-External-WithAlwaysServe.json
     */
    /**
     * Sample code: Endpoint-PUT-External-WithAlwaysServe.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointPUTExternalWithAlwaysServe(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .trafficManagerProfiles()
            .manager()
            .serviceClient()
            .getEndpoints()
            .createOrUpdateWithResponse(
                "azuresdkfornetautoresttrafficmanager1421",
                "azsmnet6386",
                EndpointTypes.EXTERNAL_ENDPOINTS,
                "azsmnet7187",
                new EndpointInner()
                    .withName("azsmnet7187")
                    .withType("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints")
                    .withTarget("foobar.contoso.com")
                    .withEndpointStatus(EndpointStatus.ENABLED)
                    .withEndpointLocation("North Europe")
                    .withAlwaysServe(AlwaysServe.ENABLED),
                com.azure.core.util.Context.NONE);
    }
}
