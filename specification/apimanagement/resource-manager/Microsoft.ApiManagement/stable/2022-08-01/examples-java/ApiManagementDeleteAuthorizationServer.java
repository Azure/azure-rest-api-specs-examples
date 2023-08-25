/** Samples for AuthorizationServer Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteAuthorizationServer.json
     */
    /**
     * Sample code: ApiManagementDeleteAuthorizationServer.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteAuthorizationServer(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizationServers()
            .deleteWithResponse("rg1", "apimService1", "newauthServer2", "*", com.azure.core.util.Context.NONE);
    }
}
