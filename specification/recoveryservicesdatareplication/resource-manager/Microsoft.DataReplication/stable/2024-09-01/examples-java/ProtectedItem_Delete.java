
/**
 * Samples for ProtectedItem Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ProtectedItem_Delete.json
     */
    /**
     * Sample code: Deletes the protected item.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void deletesTheProtectedItem(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.protectedItems().delete("rgrecoveryservicesdatareplication", "4", "d", true,
            com.azure.core.util.Context.NONE);
    }
}
