
/**
 * Samples for ApiRevision ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListApiRevisions.json
     */
    /**
     * Sample code: ApiManagementListApiRevisions.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListApiRevisions(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiRevisions().listByService("rg1", "apimService1", "57d2ef278aa04f0888cba3f3", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
