
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/getPrivateEndpointConnection.json
     */
    /**
     * Sample code: KeyVaultGetPrivateEndpointConnection.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void
        keyVaultGetPrivateEndpointConnection(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().getWithResponse("sample-group", "sample-vault",
            "sample-pec", com.azure.core.util.Context.NONE);
    }
}
