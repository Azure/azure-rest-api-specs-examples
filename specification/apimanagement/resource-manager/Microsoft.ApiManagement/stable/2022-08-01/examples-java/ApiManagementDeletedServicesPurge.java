/** Samples for DeletedServices Purge. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeletedServicesPurge.json
     */
    /**
     * Sample code: ApiManagementDeletedServicesPurge.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeletedServicesPurge(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.deletedServices().purge("apimService3", "westus", com.azure.core.util.Context.NONE);
    }
}
