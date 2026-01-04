
/**
 * Samples for Machines List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * MachineList.json
     */
    /**
     * Sample code: List Machines in an Agentpool by Managed Cluster.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listMachinesInAnAgentpoolByManagedCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getMachines().list("rg1", "clustername1", "agentpool1",
            com.azure.core.util.Context.NONE);
    }
}
