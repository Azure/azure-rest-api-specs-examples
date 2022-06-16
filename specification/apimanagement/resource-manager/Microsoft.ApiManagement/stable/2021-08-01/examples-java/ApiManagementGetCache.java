import com.azure.core.util.Context;

/** Samples for Cache Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetCache.json
     */
    /**
     * Sample code: ApiManagementGetCache.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetCache(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.caches().getWithResponse("rg1", "apimService1", "c1", Context.NONE);
    }
}
