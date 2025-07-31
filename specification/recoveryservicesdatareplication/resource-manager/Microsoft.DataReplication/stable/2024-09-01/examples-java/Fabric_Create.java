
/**
 * Samples for Fabric Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Fabric_Create.json
     */
    /**
     * Sample code: Puts the fabric.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void putsTheFabric(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.fabrics().define("wPR").withRegion((String) null).withExistingResourceGroup("rgswagger_2024-09-01")
            .create();
    }
}
