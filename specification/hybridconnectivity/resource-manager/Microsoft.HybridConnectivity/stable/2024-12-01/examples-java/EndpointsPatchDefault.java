
import com.azure.resourcemanager.hybridconnectivity.models.EndpointProperties;
import com.azure.resourcemanager.hybridconnectivity.models.EndpointResource;
import com.azure.resourcemanager.hybridconnectivity.models.Type;

/**
 * Samples for Endpoints Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/EndpointsPatchDefault.json
     */
    /**
     * Sample code: HybridConnectivityEndpointsPatchDefault.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void hybridConnectivityEndpointsPatchDefault(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        EndpointResource resource = manager.endpoints().getWithResponse(
            "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine",
            "default", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new EndpointProperties().withType(Type.DEFAULT)).apply();
    }
}
