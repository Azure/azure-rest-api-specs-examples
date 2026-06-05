
import com.azure.resourcemanager.trafficmanager.models.EndpointTypes;

/**
 * Samples for Endpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/Endpoint-GET-External-WithSubnetMapping.json
     */
    /**
     * Sample code: Endpoint-GET-External-WithSubnetMapping.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void
        endpointGETExternalWithSubnetMapping(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getEndpoints().getWithResponse("azuresdkfornetautoresttrafficmanager2191",
            "azuresdkfornetautoresttrafficmanager8224", EndpointTypes.EXTERNAL_ENDPOINTS, "My%20external%20endpoint",
            com.azure.core.util.Context.NONE);
    }
}
