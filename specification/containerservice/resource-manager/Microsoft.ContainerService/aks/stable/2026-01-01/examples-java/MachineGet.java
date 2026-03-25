
/**
 * Samples for Machines Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/MachineGet.json
     */
    /**
     * Sample code: Get a Machine in an Agent Pools by Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void getAMachineInAnAgentPoolsByManagedCluster(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMachines().getWithResponse("rg1", "clustername1", "agentpool1",
            "aks-nodepool1-42263519-vmss00000t", com.azure.core.util.Context.NONE);
    }
}
