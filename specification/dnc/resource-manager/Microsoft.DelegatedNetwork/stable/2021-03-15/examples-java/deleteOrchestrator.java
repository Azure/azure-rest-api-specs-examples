/** Samples for OrchestratorInstanceService Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/deleteOrchestrator.json
     */
    /**
     * Sample code: Delete Orchestrator Instance.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void deleteOrchestratorInstance(
        com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager.orchestratorInstanceServices().delete("TestRG", "k8stest1", null, com.azure.core.util.Context.NONE);
    }
}
