/** Samples for Authorization Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetAuthorization.json
     */
    /**
     * Sample code: ApiManagementGetAuthorization.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetAuthorization(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizations()
            .getWithResponse("rg1", "apimService1", "aadwithauthcode", "authz1", com.azure.core.util.Context.NONE);
    }
}
