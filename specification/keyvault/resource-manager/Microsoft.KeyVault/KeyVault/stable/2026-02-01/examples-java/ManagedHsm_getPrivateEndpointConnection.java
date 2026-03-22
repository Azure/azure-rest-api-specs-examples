
/**
 * Samples for MhsmPrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedHsm_getPrivateEndpointConnection.json
     */
    /**
     * Sample code: ManagedHsmGetPrivateEndpointConnection.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void
        managedHsmGetPrivateEndpointConnection(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getMhsmPrivateEndpointConnections().getWithResponse("sample-group", "sample-mhsm",
            "sample-pec", com.azure.core.util.Context.NONE);
    }
}
