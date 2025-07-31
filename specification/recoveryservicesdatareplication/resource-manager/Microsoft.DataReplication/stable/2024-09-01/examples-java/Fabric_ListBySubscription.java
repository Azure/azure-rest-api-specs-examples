
/**
 * Samples for Fabric List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Fabric_ListBySubscription.json
     */
    /**
     * Sample code: Lists the fabrics by subscription.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsTheFabricsBySubscription(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.fabrics().list(com.azure.core.util.Context.NONE);
    }
}
