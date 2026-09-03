
import com.azure.resourcemanager.containerservice.models.AgentPoolUpdate;
import com.azure.resourcemanager.containerservice.models.AgentPoolUpdateProperties;

/**
 * Samples for AgentPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/AgentPoolsUpdate_Scale.json
     */
    /**
     * Sample code: Update Agent Pool - Scale VMSS.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        updateAgentPoolScaleVMSS(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().update("rg1", "clustername1", "agentpool1",
            new AgentPoolUpdate().withProperties(new AgentPoolUpdateProperties().withCount(5)), null,
            com.azure.core.util.Context.NONE);
    }
}
