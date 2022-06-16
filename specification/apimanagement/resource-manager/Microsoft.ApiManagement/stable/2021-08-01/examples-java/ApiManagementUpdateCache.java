import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.CacheContract;

/** Samples for Cache Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateCache.json
     */
    /**
     * Sample code: ApiManagementUpdateCache.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateCache(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        CacheContract resource = manager.caches().getWithResponse("rg1", "apimService1", "c1", Context.NONE).getValue();
        resource.update().withUseFromLocation("westindia").withIfMatch("*").apply();
    }
}
