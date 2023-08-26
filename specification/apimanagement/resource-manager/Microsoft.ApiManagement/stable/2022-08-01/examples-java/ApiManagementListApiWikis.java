/** Samples for ApiWikisOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListApiWikis.json
     */
    /**
     * Sample code: ApiManagementListApiWikis.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApiWikis(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiWikisOperations()
            .list(
                "rg1", "apimService1", "57d1f7558aa04f15146d9d8a", null, null, null, com.azure.core.util.Context.NONE);
    }
}
