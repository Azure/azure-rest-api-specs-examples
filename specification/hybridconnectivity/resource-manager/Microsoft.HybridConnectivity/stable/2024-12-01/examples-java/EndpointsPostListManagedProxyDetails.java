
import com.azure.resourcemanager.hybridconnectivity.models.ManagedProxyRequest;
import com.azure.resourcemanager.hybridconnectivity.models.ServiceName;

/**
 * Samples for Endpoints ListManagedProxyDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/EndpointsPostListManagedProxyDetails.json
     */
    /**
     * Sample code: HybridConnectivityEndpointsPostListManagedProxyDetails.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void hybridConnectivityEndpointsPostListManagedProxyDetails(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.endpoints().listManagedProxyDetailsWithResponse(
            "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/arcGroup/providers/Microsoft.Compute/virtualMachines/vm00006",
            "default", new ManagedProxyRequest().withService("127.0.0.1:65035").withHostname("r.proxy.arc.com")
                .withServiceName(ServiceName.WAC),
            com.azure.core.util.Context.NONE);
    }
}
