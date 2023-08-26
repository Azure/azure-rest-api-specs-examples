/** Samples for Authorization ListByAuthorizationProvider. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListAuthorizationsAuthCode.json
     */
    /**
     * Sample code: ApiManagementListAuthorizationsAuthCode.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListAuthorizationsAuthCode(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizations()
            .listByAuthorizationProvider(
                "rg1", "apimService1", "aadwithauthcode", null, null, null, com.azure.core.util.Context.NONE);
    }
}
