
/**
 * Samples for PrivateEndpointConnections ListByResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/listPrivateEndpointConnection.json
     */
    /**
     * Sample code: KeyVaultListPrivateEndpointConnection.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void
        keyVaultListPrivateEndpointConnection(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().listByResource("sample-group", "sample-vault",
            com.azure.core.util.Context.NONE);
    }
}
