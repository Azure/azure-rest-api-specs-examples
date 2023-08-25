/** Samples for ApiDiagnostic Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteApiDiagnostic.json
     */
    /**
     * Sample code: ApiManagementDeleteApiDiagnostic.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteApiDiagnostic(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiDiagnostics()
            .deleteWithResponse(
                "rg1",
                "apimService1",
                "57d1f7558aa04f15146d9d8a",
                "applicationinsights",
                "*",
                com.azure.core.util.Context.NONE);
    }
}
