/** Samples for Diagnostic ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListDiagnostics.json
     */
    /**
     * Sample code: ApiManagementListDiagnostics.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListDiagnostics(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.diagnostics().listByService("rg1", "apimService1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
