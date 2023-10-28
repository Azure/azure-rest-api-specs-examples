/** Samples for MhsmPrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/ManagedHsm_deletePrivateEndpointConnection.json
     */
    /**
     * Sample code: ManagedHsmDeletePrivateEndpointConnection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void managedHsmDeletePrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .vaults()
            .manager()
            .serviceClient()
            .getMhsmPrivateEndpointConnections()
            .delete("sample-group", "sample-mhsm", "sample-pec", com.azure.core.util.Context.NONE);
    }
}
