
/**
 * Samples for Machines List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/MachineList.json
     */
    /**
     * Sample code: List Machines in an Agentpool by Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listMachinesInAnAgentpoolByManagedCluster(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMachines().list("rg1", "clustername1", "agentpool1",
            com.azure.core.util.Context.NONE);
    }
}
