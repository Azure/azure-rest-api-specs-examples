import com.azure.resourcemanager.attestation.models.AttestationServiceCreationSpecificParams;

/** Samples for AttestationProviders Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Create_AttestationProvider.json
     */
    /**
     * Sample code: AttestationProviders_Create.
     *
     * @param manager Entry point to AttestationManager.
     */
    public static void attestationProvidersCreate(com.azure.resourcemanager.attestation.AttestationManager manager) {
        manager
            .attestationProviders()
            .define("myattestationprovider")
            .withRegion((String) null)
            .withExistingResourceGroup("MyResourceGroup")
            .withProperties((AttestationServiceCreationSpecificParams) null)
            .create();
    }
}
