/** Samples for OrchestratorInstanceService GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/getOrchestrator.json
     */
    /**
     * Sample code: Get details of a orchestratorInstance.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void getDetailsOfAOrchestratorInstance(
        com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager
            .orchestratorInstanceServices()
            .getByResourceGroupWithResponse("TestRG", "testk8s1", com.azure.core.util.Context.NONE);
    }
}
