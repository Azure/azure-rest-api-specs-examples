import com.azure.core.util.Context;

/** Samples for ApiProduct ListByApis. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListApiProducts.json
     */
    /**
     * Sample code: ApiManagementListApiProducts.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApiProducts(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiProducts()
            .listByApis("rg1", "apimService1", "57d2ef278aa04f0888cba3f3", null, null, null, Context.NONE);
    }
}
