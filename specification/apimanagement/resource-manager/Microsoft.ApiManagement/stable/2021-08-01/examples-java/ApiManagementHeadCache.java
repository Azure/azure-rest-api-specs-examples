import com.azure.core.util.Context;

/** Samples for Cache GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadCache.json
     */
    /**
     * Sample code: ApiManagementHeadCache.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadCache(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.caches().getEntityTagWithResponse("rg1", "apimService1", "default", Context.NONE);
    }
}
