
import com.azure.resourcemanager.trafficmanager.fluent.models.EndpointInner;
import com.azure.resourcemanager.trafficmanager.models.EndpointTypes;

/**
 * Samples for Endpoints Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/Endpoint-PATCH-External-Target.json
     */
    /**
     * Sample code: Endpoint-PATCH-External-Target.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void endpointPATCHExternalTarget(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getEndpoints().updateWithResponse("azuresdkfornetautoresttrafficmanager1421",
            "azsmnet6386", EndpointTypes.EXTERNAL_ENDPOINTS, "azsmnet7187",
            new EndpointInner().withTarget("another.foobar.contoso.com"), com.azure.core.util.Context.NONE);
    }
}
