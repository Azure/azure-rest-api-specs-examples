
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Operations_List.json
     */
    /**
     * Sample code: Get a list of REST API operations supported by Microsoft.DataReplication.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void getAListOfRESTAPIOperationsSupportedByMicrosoftDataReplication(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
