/** Samples for OrchestratorInstanceService List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/orchestratorInstanceListBySub.json
     */
    /**
     * Sample code: Get orchestratorInstance resources by subscription.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void getOrchestratorInstanceResourcesBySubscription(
        com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager.orchestratorInstanceServices().list(com.azure.core.util.Context.NONE);
    }
}
