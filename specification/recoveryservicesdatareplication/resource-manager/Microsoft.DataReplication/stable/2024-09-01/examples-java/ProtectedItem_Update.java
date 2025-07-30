
import com.azure.resourcemanager.recoveryservicesdatareplication.models.ProtectedItemModel;

/**
 * Samples for ProtectedItem Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ProtectedItem_Update.json
     */
    /**
     * Sample code: Update Protected Item.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void updateProtectedItem(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        ProtectedItemModel resource = manager.protectedItems()
            .getWithResponse("rgswagger_2024-09-01", "4", "d", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
