
/**
 * Samples for Policy Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Policy_Delete.json
     */
    /**
     * Sample code: Deletes the policy.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void deletesThePolicy(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.policies().delete("rgrecoveryservicesdatareplication", "4", "wqfscsdv",
            com.azure.core.util.Context.NONE);
    }
}
