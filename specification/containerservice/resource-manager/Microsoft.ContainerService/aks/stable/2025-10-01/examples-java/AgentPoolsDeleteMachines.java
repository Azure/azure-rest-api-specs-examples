
import com.azure.resourcemanager.containerservice.models.AgentPoolDeleteMachinesParameter;
import java.util.Arrays;

/**
 * Samples for AgentPools DeleteMachines.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * AgentPoolsDeleteMachines.json
     */
    /**
     * Sample code: Delete Specific Machines in an Agent Pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteSpecificMachinesInAnAgentPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools().deleteMachines("rg1", "clustername1",
            "agentpool1",
            new AgentPoolDeleteMachinesParameter().withMachineNames(
                Arrays.asList("aks-nodepool1-42263519-vmss00000a", "aks-nodepool1-42263519-vmss00000b")),
            com.azure.core.util.Context.NONE);
    }
}
