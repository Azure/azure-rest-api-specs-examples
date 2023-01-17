/** Samples for OrchestratorInstanceService ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/orchestratorInstanceListByRG.json
     */
    /**
     * Sample code: Get OrchestratorInstance resources by resource group.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void getOrchestratorInstanceResourcesByResourceGroup(
        com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager.orchestratorInstanceServices().listByResourceGroup("testRG", com.azure.core.util.Context.NONE);
    }
}
