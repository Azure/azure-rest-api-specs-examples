import com.azure.core.util.Context;

/** Samples for AuthorizationServer ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListAuthorizationServers.json
     */
    /**
     * Sample code: ApiManagementListAuthorizationServers.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListAuthorizationServers(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.authorizationServers().listByService("rg1", "apimService1", null, null, null, Context.NONE);
    }
}
