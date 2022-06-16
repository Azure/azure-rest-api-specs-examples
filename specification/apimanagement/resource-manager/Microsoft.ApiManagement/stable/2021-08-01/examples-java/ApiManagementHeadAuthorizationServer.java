import com.azure.core.util.Context;

/** Samples for AuthorizationServer GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadAuthorizationServer.json
     */
    /**
     * Sample code: ApiManagementHeadAuthorizationServer.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadAuthorizationServer(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.authorizationServers().getEntityTagWithResponse("rg1", "apimService1", "newauthServer2", Context.NONE);
    }
}
