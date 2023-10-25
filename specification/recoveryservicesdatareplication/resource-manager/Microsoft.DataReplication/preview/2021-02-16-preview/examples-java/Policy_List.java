/** Samples for Policy List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Policy_List.json
     */
    /**
     * Sample code: Policy_List.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void policyList(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.policies().list("rgrecoveryservicesdatareplication", "4", com.azure.core.util.Context.NONE);
    }
}
