/** Samples for AttestationProviders ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Get_AttestationProvidersListByResourceGroup.json
     */
    /**
     * Sample code: AttestationProviders_ListByResourceGroup.
     *
     * @param manager Entry point to AttestationManager.
     */
    public static void attestationProvidersListByResourceGroup(
        com.azure.resourcemanager.attestation.AttestationManager manager) {
        manager.attestationProviders().listByResourceGroupWithResponse("testrg1", com.azure.core.util.Context.NONE);
    }
}
