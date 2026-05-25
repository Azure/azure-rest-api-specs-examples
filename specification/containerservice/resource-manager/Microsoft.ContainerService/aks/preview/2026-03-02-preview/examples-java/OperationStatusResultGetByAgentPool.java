
/**
 * Samples for OperationStatusResult GetByAgentPool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/OperationStatusResultGetByAgentPool.json
     */
    /**
     * Sample code: Get OperationStatusResult.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getOperationStatusResult(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getOperationStatusResults().getByAgentPoolWithResponse("rg1", "clustername1",
            "agentpool1", "00000000-0000-0000-0000-000000000001", com.azure.core.util.Context.NONE);
    }
}
