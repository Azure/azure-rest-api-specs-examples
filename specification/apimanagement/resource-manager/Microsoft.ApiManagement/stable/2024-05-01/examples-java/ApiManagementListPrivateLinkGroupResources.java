
/**
 * Samples for PrivateEndpointConnection ListPrivateLinkResources.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListPrivateLinkGroupResources.json
     */
    /**
     * Sample code: ApiManagementListPrivateLinkGroupResources.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListPrivateLinkGroupResources(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.privateEndpointConnections().listPrivateLinkResourcesWithResponse("rg1", "apimService1",
            com.azure.core.util.Context.NONE);
    }
}
