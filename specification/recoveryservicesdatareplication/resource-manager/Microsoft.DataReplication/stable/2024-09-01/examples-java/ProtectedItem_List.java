
/**
 * Samples for ProtectedItem List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ProtectedItem_List.json
     */
    /**
     * Sample code: Lists the protected items.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsTheProtectedItems(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.protectedItems().list("rgrecoveryservicesdatareplication", "4", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
