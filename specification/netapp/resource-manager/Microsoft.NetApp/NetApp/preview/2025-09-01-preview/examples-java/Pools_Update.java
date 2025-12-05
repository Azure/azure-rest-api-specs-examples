
import com.azure.resourcemanager.netapp.models.CapacityPool;

/**
 * Samples for Pools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Pools_Update.json
     */
    /**
     * Sample code: Pools_Update.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void poolsUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        CapacityPool resource
            = manager.pools().getWithResponse("myRG", "account1", "pool1", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
