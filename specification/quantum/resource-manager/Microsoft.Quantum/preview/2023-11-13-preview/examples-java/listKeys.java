
/**
 * Samples for WorkspaceOperation ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quantum/resource-manager/Microsoft.Quantum/preview/2023-11-13-preview/examples/listKeys.json
     */
    /**
     * Sample code: ListKeys.
     * 
     * @param manager Entry point to AzureQuantumManager.
     */
    public static void listKeys(com.azure.resourcemanager.quantum.AzureQuantumManager manager) {
        manager.workspaceOperations().listKeysWithResponse("quantumResourcegroup", "quantumworkspace1",
            com.azure.core.util.Context.NONE);
    }
}
