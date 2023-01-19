/** Samples for Clusters ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterListByResourceGroupOperation_example.json
     */
    /**
     * Sample code: List cluster by resource group.
     *
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void listClusterByResourceGroup(
        com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.clusters().listByResourceGroupWithResponse("resRg", com.azure.core.util.Context.NONE);
    }
}
