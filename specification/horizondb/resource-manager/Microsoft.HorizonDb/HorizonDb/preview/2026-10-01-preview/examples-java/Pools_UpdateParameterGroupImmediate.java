
import com.azure.resourcemanager.horizondb.models.ParameterApplyMethod;
import com.azure.resourcemanager.horizondb.models.Pool;
import com.azure.resourcemanager.horizondb.models.PoolParameterGroupReferenceForPatchUpdate;
import com.azure.resourcemanager.horizondb.models.PoolPropertiesForPatchUpdate;

/**
 * Samples for Pools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Pools_UpdateParameterGroupImmediate.json
     */
    /**
     * Sample code: Associate a HorizonDB pool with a parameter group immediately.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void associateAHorizonDBPoolWithAParameterGroupImmediately(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        Pool resource = manager.pools().getWithResponse("example-resource-group", "example-cluster", "example-pool-01",
            com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new PoolPropertiesForPatchUpdate()
            .withParameterGroup(new PoolParameterGroupReferenceForPatchUpdate().withParameterGroupResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/parameterGroups/production-parameters")
                .withApplyMethod(ParameterApplyMethod.IMMEDIATE)))
            .apply();
    }
}
