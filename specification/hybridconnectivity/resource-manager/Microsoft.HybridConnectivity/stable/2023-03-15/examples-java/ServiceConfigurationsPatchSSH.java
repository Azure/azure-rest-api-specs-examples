
import com.azure.resourcemanager.hybridconnectivity.models.ServiceConfigurationResource;

/**
 * Samples for ServiceConfigurations Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/
     * ServiceConfigurationsPatchSSH.json
     */
    /**
     * Sample code: ServiceConfigurationsPatchSSH.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void
        serviceConfigurationsPatchSSH(com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        ServiceConfigurationResource resource = manager.serviceConfigurations().getWithResponse(
            "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default",
            "default", "SSH", com.azure.core.util.Context.NONE).getValue();
        resource.update().withPort(22L).apply();
    }
}
