/** Samples for AuthorizationServer ListSecrets. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementAuthorizationServerListSecrets.json
     */
    /**
     * Sample code: ApiManagementAuthorizationServerListSecrets.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementAuthorizationServerListSecrets(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizationServers()
            .listSecretsWithResponse("rg1", "apimService1", "newauthServer2", com.azure.core.util.Context.NONE);
    }
}
