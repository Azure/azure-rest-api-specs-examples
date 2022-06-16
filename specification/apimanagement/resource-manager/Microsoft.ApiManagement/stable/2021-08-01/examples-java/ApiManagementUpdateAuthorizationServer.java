import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.AuthorizationServerContract;

/** Samples for AuthorizationServer Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateAuthorizationServer.json
     */
    /**
     * Sample code: ApiManagementUpdateAuthorizationServer.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateAuthorizationServer(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        AuthorizationServerContract resource =
            manager
                .authorizationServers()
                .getWithResponse("rg1", "apimService1", "newauthServer", Context.NONE)
                .getValue();
        resource.update().withClientId("update").withClientSecret("updated").withIfMatch("*").apply();
    }
}
