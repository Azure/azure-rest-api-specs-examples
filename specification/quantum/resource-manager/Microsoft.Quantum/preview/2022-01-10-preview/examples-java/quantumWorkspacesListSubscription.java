/** Samples for Workspaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/quantum/resource-manager/Microsoft.Quantum/preview/2022-01-10-preview/examples/quantumWorkspacesListSubscription.json
     */
    /**
     * Sample code: QuantumWorkspacesListBySubscription.
     *
     * @param manager Entry point to AzureQuantumManager.
     */
    public static void quantumWorkspacesListBySubscription(
        com.azure.resourcemanager.quantum.AzureQuantumManager manager) {
        manager.workspaces().list(com.azure.core.util.Context.NONE);
    }
}
