/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to AttestationManager.
     */
    public static void operationsList(com.azure.resourcemanager.attestation.AttestationManager manager) {
        manager.operations().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
