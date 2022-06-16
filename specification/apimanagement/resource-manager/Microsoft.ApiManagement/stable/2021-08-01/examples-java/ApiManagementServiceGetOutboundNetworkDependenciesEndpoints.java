import com.azure.core.util.Context;

/** Samples for OutboundNetworkDependenciesEndpoints ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementServiceGetOutboundNetworkDependenciesEndpoints.json
     */
    /**
     * Sample code: ApiManagementServiceGetOutboundNetworkDependenciesEndpoints.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementServiceGetOutboundNetworkDependenciesEndpoints(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.outboundNetworkDependenciesEndpoints().listByServiceWithResponse("rg1", "apimService1", Context.NONE);
    }
}
