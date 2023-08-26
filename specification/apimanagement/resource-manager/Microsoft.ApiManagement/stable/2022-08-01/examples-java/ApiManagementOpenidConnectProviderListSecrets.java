/** Samples for OpenIdConnectProvider ListSecrets. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementOpenidConnectProviderListSecrets.json
     */
    /**
     * Sample code: ApiManagementOpenidConnectProviderListSecrets.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementOpenidConnectProviderListSecrets(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .openIdConnectProviders()
            .listSecretsWithResponse("rg1", "apimService1", "templateOpenIdConnect2", com.azure.core.util.Context.NONE);
    }
}
