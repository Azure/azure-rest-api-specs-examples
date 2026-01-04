
/**
 * Samples for MhsmPrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ManagedHsm_getPrivateEndpointConnection.json
     */
    /**
     * Sample code: ManagedHsmGetPrivateEndpointConnection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void managedHsmGetPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getMhsmPrivateEndpointConnections().getWithResponse("sample-group",
            "sample-mhsm", "sample-pec", com.azure.core.util.Context.NONE);
    }
}
