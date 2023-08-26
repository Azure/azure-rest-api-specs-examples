/** Samples for Tag ListByApi. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListApiTags.json
     */
    /**
     * Sample code: ApiManagementListApiTags.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApiTags(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tags()
            .listByApi(
                "rg1", "apimService1", "57d2ef278aa04f0888cba3f3", null, null, null, com.azure.core.util.Context.NONE);
    }
}
