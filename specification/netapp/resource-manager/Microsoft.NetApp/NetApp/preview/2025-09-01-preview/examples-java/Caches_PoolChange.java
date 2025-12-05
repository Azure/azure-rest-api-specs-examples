
import com.azure.resourcemanager.netapp.models.PoolChangeRequest;

/**
 * Samples for Caches PoolChange.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Caches_PoolChange.json
     */
    /**
     * Sample code: Caches_PoolChange.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void cachesPoolChange(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.caches().poolChange("myRG", "account1", "pool1", "cache1",
            new PoolChangeRequest().withNewPoolResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool2"),
            com.azure.core.util.Context.NONE);
    }
}
