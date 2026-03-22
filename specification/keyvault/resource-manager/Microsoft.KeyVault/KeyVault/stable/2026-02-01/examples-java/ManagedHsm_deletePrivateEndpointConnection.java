
/**
 * Samples for MhsmPrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedHsm_deletePrivateEndpointConnection.json
     */
    /**
     * Sample code: ManagedHsmDeletePrivateEndpointConnection.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void
        managedHsmDeletePrivateEndpointConnection(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getMhsmPrivateEndpointConnections().delete("sample-group", "sample-mhsm", "sample-pec",
            com.azure.core.util.Context.NONE);
    }
}
