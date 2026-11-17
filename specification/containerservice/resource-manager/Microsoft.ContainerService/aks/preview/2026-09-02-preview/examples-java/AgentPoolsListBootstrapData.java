
/**
 * Samples for AgentPools ListBootstrapData.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-02-preview/AgentPoolsListBootstrapData.json
     */
    /**
     * Sample code: List Bootstrap Data for FlexNode Agent Pool.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listBootstrapDataForFlexNodeAgentPool(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().listBootstrapDataWithResponse("rg1", "clustername1", "flexnode1",
            com.azure.core.util.Context.NONE);
    }
}
