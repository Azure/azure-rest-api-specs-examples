
/**
 * Samples for Tag ListByOperation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/
     * ApiManagementListApiOperationTags.json
     */
    /**
     * Sample code: ApiManagementListApiOperationTags.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListApiOperationTags(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tags().listByOperation("rg1", "apimService1", "57d2ef278aa04f0888cba3f3", "57d2ef278aa04f0888cba3f6",
            null, null, null, com.azure.core.util.Context.NONE);
    }
}
