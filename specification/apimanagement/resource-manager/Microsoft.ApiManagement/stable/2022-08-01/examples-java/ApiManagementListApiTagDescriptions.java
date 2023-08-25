/** Samples for ApiTagDescription ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListApiTagDescriptions.json
     */
    /**
     * Sample code: ApiManagementListApiTagDescriptions.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApiTagDescriptions(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiTagDescriptions()
            .listByService(
                "rg1", "apimService1", "57d2ef278aa04f0888cba3f3", null, null, null, com.azure.core.util.Context.NONE);
    }
}
