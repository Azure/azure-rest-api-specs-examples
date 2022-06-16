import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnection GetPrivateLinkResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetPrivateLinkGroupResource.json
     */
    /**
     * Sample code: ApiManagementGetPrivateLinkGroupResource.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetPrivateLinkGroupResource(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .privateEndpointConnections()
            .getPrivateLinkResourceWithResponse("rg1", "apimService1", "privateLinkSubResourceName", Context.NONE);
    }
}
