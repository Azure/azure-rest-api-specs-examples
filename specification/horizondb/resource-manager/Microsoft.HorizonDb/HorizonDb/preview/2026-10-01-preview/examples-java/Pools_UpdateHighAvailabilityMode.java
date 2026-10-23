
import com.azure.resourcemanager.horizondb.models.Pool;
import com.azure.resourcemanager.horizondb.models.PoolHighAvailabilityMode;
import com.azure.resourcemanager.horizondb.models.PoolHighAvailabilityPatch;
import com.azure.resourcemanager.horizondb.models.PoolPropertiesForPatchUpdate;

/**
 * Samples for Pools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Pools_UpdateHighAvailabilityMode.json
     */
    /**
     * Sample code: Require zone-redundant placement for a HorizonDB pool.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        requireZoneRedundantPlacementForAHorizonDBPool(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        Pool resource = manager.pools().getWithResponse("example-resource-group", "example-cluster", "example-pool-01",
            com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new PoolPropertiesForPatchUpdate()
            .withHighAvailability(new PoolHighAvailabilityPatch().withMode(PoolHighAvailabilityMode.ZONE_REDUNDANT)))
            .apply();
    }
}
