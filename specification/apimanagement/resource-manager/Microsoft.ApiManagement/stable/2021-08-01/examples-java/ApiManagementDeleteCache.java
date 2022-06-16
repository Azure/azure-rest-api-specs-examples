import com.azure.core.util.Context;

/** Samples for Cache Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteCache.json
     */
    /**
     * Sample code: ApiManagementDeleteCache.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteCache(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.caches().deleteWithResponse("rg1", "apimService1", "southindia", "*", Context.NONE);
    }
}
