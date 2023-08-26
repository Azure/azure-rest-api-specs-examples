/** Samples for ApiManagementService GetSsoToken. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceGetSsoToken.json
     */
    /**
     * Sample code: ApiManagementServiceGetSsoToken.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementServiceGetSsoToken(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiManagementServices()
            .getSsoTokenWithResponse("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
