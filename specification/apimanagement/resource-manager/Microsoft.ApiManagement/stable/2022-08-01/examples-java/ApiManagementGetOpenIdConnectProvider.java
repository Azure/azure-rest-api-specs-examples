/** Samples for OpenIdConnectProvider Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetOpenIdConnectProvider.json
     */
    /**
     * Sample code: ApiManagementGetOpenIdConnectProvider.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetOpenIdConnectProvider(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .openIdConnectProviders()
            .getWithResponse("rg1", "apimService1", "templateOpenIdConnect2", com.azure.core.util.Context.NONE);
    }
}
