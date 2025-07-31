
/**
 * Samples for Policy List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Policy_List.json
     */
    /**
     * Sample code: Lists the policies.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsThePolicies(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.policies().list("rgrecoveryservicesdatareplication", "4", com.azure.core.util.Context.NONE);
    }
}
