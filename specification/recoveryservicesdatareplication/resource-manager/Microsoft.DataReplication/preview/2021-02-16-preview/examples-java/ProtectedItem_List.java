/** Samples for ProtectedItem List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ProtectedItem_List.json
     */
    /**
     * Sample code: ProtectedItem_List.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void protectedItemList(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.protectedItems().list("rgrecoveryservicesdatareplication", "4", com.azure.core.util.Context.NONE);
    }
}
