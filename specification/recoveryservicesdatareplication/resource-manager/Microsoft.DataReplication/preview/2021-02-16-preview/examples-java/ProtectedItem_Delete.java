/** Samples for ProtectedItem Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ProtectedItem_Delete.json
     */
    /**
     * Sample code: ProtectedItem_Delete.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void protectedItemDelete(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .protectedItems()
            .delete("rgrecoveryservicesdatareplication", "4", "d", true, com.azure.core.util.Context.NONE);
    }
}
