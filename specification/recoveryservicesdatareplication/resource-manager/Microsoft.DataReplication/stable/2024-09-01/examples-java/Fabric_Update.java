
import com.azure.resourcemanager.recoveryservicesdatareplication.models.FabricModel;

/**
 * Samples for Fabric Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Fabric_Update.json
     */
    /**
     * Sample code: Updates the fabric.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void updatesTheFabric(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        FabricModel resource = manager.fabrics()
            .getByResourceGroupWithResponse("rgswagger_2024-09-01", "wPR", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
