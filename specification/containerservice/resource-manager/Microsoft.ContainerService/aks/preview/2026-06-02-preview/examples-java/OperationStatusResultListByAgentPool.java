
/**
 * Samples for OperationStatusResult ListByAgentPool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/OperationStatusResultListByAgentPool.json
     */
    /**
     * Sample code: List Operations on Agent Pool.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listOperationsOnAgentPool(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getOperationStatusResults().listByAgentPool("rg1", "clustername1", "agentpool1", null,
            com.azure.core.util.Context.NONE);
    }
}
