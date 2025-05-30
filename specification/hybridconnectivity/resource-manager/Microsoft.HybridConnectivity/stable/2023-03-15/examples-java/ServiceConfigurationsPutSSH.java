
import com.azure.resourcemanager.hybridconnectivity.models.ServiceName;

/**
 * Samples for ServiceConfigurations CreateOrupdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/
     * ServiceConfigurationsPutSSH.json
     */
    /**
     * Sample code: ServiceConfigurationsPutSSH.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void
        serviceConfigurationsPutSSH(com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.serviceConfigurations().define("SSH").withExistingEndpoint(
            "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default",
            "default").withServiceName(ServiceName.SSH).withPort(22L).create();
    }
}
