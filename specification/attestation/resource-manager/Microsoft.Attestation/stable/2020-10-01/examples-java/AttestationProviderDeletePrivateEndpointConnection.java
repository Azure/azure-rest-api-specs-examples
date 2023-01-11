/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/AttestationProviderDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: AttestationProviderDeletePrivateEndpointConnection.
     *
     * @param manager Entry point to AttestationManager.
     */
    public static void attestationProviderDeletePrivateEndpointConnection(
        com.azure.resourcemanager.attestation.AttestationManager manager) {
        manager
            .privateEndpointConnections()
            .deleteWithResponse(
                "res6977", "sto2527", "{privateEndpointConnectionName}", com.azure.core.util.Context.NONE);
    }
}
