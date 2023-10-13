/** Samples for PrivateEndpoint GetOperationStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/PrivateEndpointConnection/GetPrivateEndpointConnectionOperationStatus.json
     */
    /**
     * Sample code: Get OperationStatus.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getOperationStatus(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .privateEndpoints()
            .getOperationStatusWithResponse(
                "gaallavaultbvtd2msi",
                "gaallaRG",
                "gaallatestpe2.5704c932-249a-490b-a142-1396838cd3b",
                "0f48183b-0a44-4dca-aec1-bba5daab888a",
                com.azure.core.util.Context.NONE);
    }
}
