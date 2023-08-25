/** Samples for Diagnostic GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadDiagnostic.json
     */
    /**
     * Sample code: ApiManagementHeadDiagnostic.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadDiagnostic(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .diagnostics()
            .getEntityTagWithResponse("rg1", "apimService1", "applicationinsights", com.azure.core.util.Context.NONE);
    }
}
