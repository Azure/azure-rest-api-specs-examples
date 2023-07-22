/** Samples for Workspaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/quantum/resource-manager/Microsoft.Quantum/preview/2022-01-10-preview/examples/quantumWorkspacesDelete.json
     */
    /**
     * Sample code: QuantumWorkspacesDelete.
     *
     * @param manager Entry point to AzureQuantumManager.
     */
    public static void quantumWorkspacesDelete(com.azure.resourcemanager.quantum.AzureQuantumManager manager) {
        manager.workspaces().delete("quantumResourcegroup", "quantumworkspace1", com.azure.core.util.Context.NONE);
    }
}
