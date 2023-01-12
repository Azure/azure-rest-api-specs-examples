/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/AttestationProviderGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: AttestationProviderGetPrivateEndpointConnection.
     *
     * @param manager Entry point to AttestationManager.
     */
    public static void attestationProviderGetPrivateEndpointConnection(
        com.azure.resourcemanager.attestation.AttestationManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse("res6977", "sto2527", "{privateEndpointConnectionName}", com.azure.core.util.Context.NONE);
    }
}
