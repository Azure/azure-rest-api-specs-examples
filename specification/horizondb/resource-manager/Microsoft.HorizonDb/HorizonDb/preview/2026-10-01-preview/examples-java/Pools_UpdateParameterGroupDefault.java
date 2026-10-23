
import com.azure.resourcemanager.horizondb.models.ParameterApplyMethod;
import com.azure.resourcemanager.horizondb.models.Pool;
import com.azure.resourcemanager.horizondb.models.PoolParameterGroupReferenceForPatchUpdate;
import com.azure.resourcemanager.horizondb.models.PoolPropertiesForPatchUpdate;

/**
 * Samples for Pools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Pools_UpdateParameterGroupDefault.json
     */
    /**
     * Sample code: Switch a HorizonDB pool back to its corresponding default parameter group.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void switchAHorizonDBPoolBackToItsCorrespondingDefaultParameterGroup(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        Pool resource = manager.pools().getWithResponse("example-resource-group", "example-cluster", "example-pool-01",
            com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new PoolPropertiesForPatchUpdate()
            .withParameterGroup(new PoolParameterGroupReferenceForPatchUpdate().withParameterGroupResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/providers/Microsoft.HorizonDb/locations/eastus2/parameterGroups/default-postgresql-17-provisioned")
                .withApplyMethod(ParameterApplyMethod.IMMEDIATE)))
            .apply();
    }
}
