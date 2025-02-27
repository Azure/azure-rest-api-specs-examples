
import com.azure.resourcemanager.hybridconnectivity.models.ServiceName;

/**
 * Samples for ServiceConfigurations CreateOrupdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/ServiceConfigurationsPutWAC.json
     */
    /**
     * Sample code: ServiceConfigurationsPutWAC.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void
        serviceConfigurationsPutWAC(com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.serviceConfigurations().define("WAC").withExistingEndpoint(
            "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default",
            "default").withServiceName(ServiceName.WAC).withPort(6516L).create();
    }
}
