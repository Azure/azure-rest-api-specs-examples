/** Samples for AttestationProviders List. */
public final class Main {
    /*
     * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Get_AttestationProvidersList.json
     */
    /**
     * Sample code: AttestationProviders_List.
     *
     * @param manager Entry point to AttestationManager.
     */
    public static void attestationProvidersList(com.azure.resourcemanager.attestation.AttestationManager manager) {
        manager.attestationProviders().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
