/** Samples for AttestationProviders ListDefault. */
public final class Main {
    /*
     * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Get_DefaultProviders.json
     */
    /**
     * Sample code: AttestationProviders_GetDefault.
     *
     * @param manager Entry point to AttestationManager.
     */
    public static void attestationProvidersGetDefault(
        com.azure.resourcemanager.attestation.AttestationManager manager) {
        manager.attestationProviders().listDefaultWithResponse(com.azure.core.util.Context.NONE);
    }
}
