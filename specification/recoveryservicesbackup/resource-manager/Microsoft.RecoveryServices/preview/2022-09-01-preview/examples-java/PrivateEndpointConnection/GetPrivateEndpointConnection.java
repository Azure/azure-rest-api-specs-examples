import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnection Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/PrivateEndpointConnection/GetPrivateEndpointConnection.json
     */
    /**
     * Sample code: Get PrivateEndpointConnection.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getPrivateEndpointConnection(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse(
                "gaallavaultbvtd2msi", "gaallaRG", "gaallatestpe2.5704c932-249a-490b-a142-1396838cd3b", Context.NONE);
    }
}
