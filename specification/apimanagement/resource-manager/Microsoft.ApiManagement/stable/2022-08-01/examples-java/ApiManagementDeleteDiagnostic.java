/** Samples for Diagnostic Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteDiagnostic.json
     */
    /**
     * Sample code: ApiManagementDeleteDiagnostic.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteDiagnostic(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .diagnostics()
            .deleteWithResponse("rg1", "apimService1", "applicationinsights", "*", com.azure.core.util.Context.NONE);
    }
}
