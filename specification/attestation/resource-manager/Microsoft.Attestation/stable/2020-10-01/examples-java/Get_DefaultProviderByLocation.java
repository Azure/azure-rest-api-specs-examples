/** Samples for AttestationProviders GetDefaultByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Get_DefaultProviderByLocation.json
     */
    /**
     * Sample code: AttestationProviders_GetDefaultWithLocation.
     *
     * @param manager Entry point to AttestationManager.
     */
    public static void attestationProvidersGetDefaultWithLocation(
        com.azure.resourcemanager.attestation.AttestationManager manager) {
        manager.attestationProviders().getDefaultByLocationWithResponse("Central US", com.azure.core.util.Context.NONE);
    }
}
