/** Samples for AuthorizationServer Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetAuthorizationServer.json
     */
    /**
     * Sample code: ApiManagementGetAuthorizationServer.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetAuthorizationServer(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizationServers()
            .getWithResponse("rg1", "apimService1", "newauthServer2", com.azure.core.util.Context.NONE);
    }
}
