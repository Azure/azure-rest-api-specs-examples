
/**
 * Samples for ApiDiagnostic ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListApiDiagnostics.json
     */
    /**
     * Sample code: ApiManagementListApiDiagnostics.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListApiDiagnostics(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiDiagnostics().listByService("rg1", "apimService1", "echo-api", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
