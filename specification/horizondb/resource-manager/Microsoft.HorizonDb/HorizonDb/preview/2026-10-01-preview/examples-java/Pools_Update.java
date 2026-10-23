
import com.azure.resourcemanager.horizondb.models.Pool;
import com.azure.resourcemanager.horizondb.models.PoolComputeModelForPatchUpdate;
import com.azure.resourcemanager.horizondb.models.PoolPropertiesForPatchUpdate;
import com.azure.resourcemanager.horizondb.models.ScheduledMaintenanceManagementType;
import com.azure.resourcemanager.horizondb.models.ScheduledMaintenancePropertiesForPatchUpdate;

/**
 * Samples for Pools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Pools_Update.json
     */
    /**
     * Sample code: Update a single HorizonDB pool compute property.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        updateASingleHorizonDBPoolComputeProperty(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        Pool resource = manager.pools().getWithResponse("example-resource-group", "example-cluster", "example-pool-01",
            com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(
                new PoolPropertiesForPatchUpdate().withComputeModel(new PoolComputeModelForPatchUpdate().withVCores(16))
                    .withScheduledMaintenance(new ScheduledMaintenancePropertiesForPatchUpdate()
                        .withMaintenanceManagementType(ScheduledMaintenanceManagementType.SYSTEM)))
            .apply();
    }
}
