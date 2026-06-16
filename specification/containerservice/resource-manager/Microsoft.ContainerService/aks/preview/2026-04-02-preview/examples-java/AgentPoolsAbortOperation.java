
/**
 * Samples for AgentPools AbortLatestOperation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/AgentPoolsAbortOperation.json
     */
    /**
     * Sample code: Abort operation on agent pool.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        abortOperationOnAgentPool(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().abortLatestOperation("rg1", "clustername1", "agentpool1",
            com.azure.core.util.Context.NONE);
    }
}
