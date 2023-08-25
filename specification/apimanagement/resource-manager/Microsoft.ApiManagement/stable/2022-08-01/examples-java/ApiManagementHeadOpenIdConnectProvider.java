/** Samples for OpenIdConnectProvider GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadOpenIdConnectProvider.json
     */
    /**
     * Sample code: ApiManagementHeadOpenIdConnectProvider.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadOpenIdConnectProvider(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .openIdConnectProviders()
            .getEntityTagWithResponse(
                "rg1", "apimService1", "templateOpenIdConnect2", com.azure.core.util.Context.NONE);
    }
}
