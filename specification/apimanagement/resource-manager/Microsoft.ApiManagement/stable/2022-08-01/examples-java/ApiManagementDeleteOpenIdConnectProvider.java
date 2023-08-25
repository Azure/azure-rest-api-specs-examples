/** Samples for OpenIdConnectProvider Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteOpenIdConnectProvider.json
     */
    /**
     * Sample code: ApiManagementDeleteOpenIdConnectProvider.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteOpenIdConnectProvider(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .openIdConnectProviders()
            .deleteWithResponse("rg1", "apimService1", "templateOpenIdConnect3", "*", com.azure.core.util.Context.NONE);
    }
}
