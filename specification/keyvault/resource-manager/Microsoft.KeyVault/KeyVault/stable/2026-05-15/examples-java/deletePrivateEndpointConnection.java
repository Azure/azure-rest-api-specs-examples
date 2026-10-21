
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15/deletePrivateEndpointConnection.json
     */
    /**
     * Sample code: KeyVaultDeletePrivateEndpointConnection.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void
        keyVaultDeletePrivateEndpointConnection(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().delete("sample-group", "sample-vault", "sample-pec",
            com.azure.core.util.Context.NONE);
    }
}
