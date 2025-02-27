
import com.azure.resourcemanager.hybridconnectivity.models.ListCredentialsRequest;
import com.azure.resourcemanager.hybridconnectivity.models.ServiceName;

/**
 * Samples for Endpoints ListCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/EndpointsPostListCredentials.json
     */
    /**
     * Sample code: HybridConnectivityEndpointsPostListCredentials.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void hybridConnectivityEndpointsPostListCredentials(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.endpoints().listCredentialsWithResponse(
            "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine",
            "default", 10800L, new ListCredentialsRequest().withServiceName(ServiceName.SSH),
            com.azure.core.util.Context.NONE);
    }
}
