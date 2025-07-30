
/**
 * Samples for Fabric ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Fabric_List.json
     */
    /**
     * Sample code: Lists the fabrics.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsTheFabrics(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.fabrics().listByResourceGroup("rgswagger_2024-09-01", "jw", com.azure.core.util.Context.NONE);
    }
}
