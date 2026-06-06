
import com.azure.resourcemanager.trafficmanager.models.EndpointTypes;

/**
 * Samples for Endpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/Endpoint-GET-External-WithLocation.json
     */
    /**
     * Sample code: Endpoint-GET-External-WithLocation.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void
        endpointGETExternalWithLocation(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getEndpoints().getWithResponse("azuresdkfornetautoresttrafficmanager1421",
            "azsmnet6386", EndpointTypes.EXTERNAL_ENDPOINTS, "azsmnet7187", com.azure.core.util.Context.NONE);
    }
}
