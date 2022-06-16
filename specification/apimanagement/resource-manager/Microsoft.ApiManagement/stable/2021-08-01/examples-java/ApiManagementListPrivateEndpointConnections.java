import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnection ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListPrivateEndpointConnections.json
     */
    /**
     * Sample code: ApiManagementListPrivateEndpointConnections.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListPrivateEndpointConnections(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.privateEndpointConnections().listByService("rg1", "apimService1", Context.NONE);
    }
}
