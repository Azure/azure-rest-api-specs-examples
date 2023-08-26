/** Samples for Authorization ListByAuthorizationProvider. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListAuthorizationsClientCred.json
     */
    /**
     * Sample code: ApiManagementListAuthorizationsClientCred.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListAuthorizationsClientCred(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizations()
            .listByAuthorizationProvider(
                "rg1", "apimService1", "aadwithclientcred", null, null, null, com.azure.core.util.Context.NONE);
    }
}
