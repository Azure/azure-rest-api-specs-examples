/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/AttestationProviderListPrivateEndpointConnections.json
     */
    /**
     * Sample code: AttestationProviderListPrivateEndpointConnections.
     *
     * @param manager Entry point to AttestationManager.
     */
    public static void attestationProviderListPrivateEndpointConnections(
        com.azure.resourcemanager.attestation.AttestationManager manager) {
        manager.privateEndpointConnections().list("res6977", "sto2527", com.azure.core.util.Context.NONE);
    }
}
