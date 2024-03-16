
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quantum/resource-manager/Microsoft.Quantum/preview/2023-11-13-preview/examples/operations.json
     */
    /**
     * Sample code: Operations.
     * 
     * @param manager Entry point to AzureQuantumManager.
     */
    public static void operations(com.azure.resourcemanager.quantum.AzureQuantumManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
