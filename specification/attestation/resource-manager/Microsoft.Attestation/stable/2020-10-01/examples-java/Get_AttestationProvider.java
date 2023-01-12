/** Samples for AttestationProviders GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Get_AttestationProvider.json
     */
    /**
     * Sample code: AttestationProviders_Get.
     *
     * @param manager Entry point to AttestationManager.
     */
    public static void attestationProvidersGet(com.azure.resourcemanager.attestation.AttestationManager manager) {
        manager
            .attestationProviders()
            .getByResourceGroupWithResponse(
                "MyResourceGroup", "myattestationprovider", com.azure.core.util.Context.NONE);
    }
}
