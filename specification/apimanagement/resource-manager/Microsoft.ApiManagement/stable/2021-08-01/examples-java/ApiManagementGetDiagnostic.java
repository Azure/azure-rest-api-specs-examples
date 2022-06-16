import com.azure.core.util.Context;

/** Samples for Diagnostic Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetDiagnostic.json
     */
    /**
     * Sample code: ApiManagementGetDiagnostic.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetDiagnostic(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.diagnostics().getWithResponse("rg1", "apimService1", "applicationinsights", Context.NONE);
    }
}
