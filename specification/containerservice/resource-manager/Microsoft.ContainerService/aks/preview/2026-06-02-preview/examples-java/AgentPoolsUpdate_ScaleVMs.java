
import com.azure.resourcemanager.containerservice.models.AgentPoolUpdate;
import com.azure.resourcemanager.containerservice.models.AgentPoolUpdateManualScaleProfile;
import com.azure.resourcemanager.containerservice.models.AgentPoolUpdateProperties;
import com.azure.resourcemanager.containerservice.models.AgentPoolUpdateScaleProfile;
import com.azure.resourcemanager.containerservice.models.AgentPoolUpdateVirtualMachinesProfile;
import java.util.Arrays;

/**
 * Samples for AgentPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/AgentPoolsUpdate_ScaleVMs.json
     */
    /**
     * Sample code: Update Agent Pool - Scale VirtualMachines.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void updateAgentPoolScaleVirtualMachines(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().update("rg1", "clustername1", "agentpool1",
            new AgentPoolUpdate().withProperties(new AgentPoolUpdateProperties().withVirtualMachinesProfile(
                new AgentPoolUpdateVirtualMachinesProfile().withScale(new AgentPoolUpdateScaleProfile().withManual(
                    Arrays.asList(new AgentPoolUpdateManualScaleProfile().withSize("Standard_D2_v2").withCount(5),
                        new AgentPoolUpdateManualScaleProfile().withSize("Standard_D2_v3").withCount(3)))))),
            null, com.azure.core.util.Context.NONE);
    }
}
