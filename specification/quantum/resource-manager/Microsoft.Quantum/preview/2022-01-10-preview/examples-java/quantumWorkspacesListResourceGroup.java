/** Samples for Workspaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/quantum/resource-manager/Microsoft.Quantum/preview/2022-01-10-preview/examples/quantumWorkspacesListResourceGroup.json
     */
    /**
     * Sample code: QuantumWorkspacesListByResourceGroup.
     *
     * @param manager Entry point to AzureQuantumManager.
     */
    public static void quantumWorkspacesListByResourceGroup(
        com.azure.resourcemanager.quantum.AzureQuantumManager manager) {
        manager.workspaces().listByResourceGroup("quantumResourcegroup", com.azure.core.util.Context.NONE);
    }
}
