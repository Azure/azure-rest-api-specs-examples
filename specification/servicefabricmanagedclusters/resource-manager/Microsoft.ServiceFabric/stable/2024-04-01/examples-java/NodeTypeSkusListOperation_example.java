
/**
 * Samples for NodeTypeSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/
     * NodeTypeSkusListOperation_example.json
     */
    /**
     * Sample code: List a node type SKUs.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void listANodeTypeSKUs(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypeSkus().list("resRg", "myCluster", "BE", com.azure.core.util.Context.NONE);
    }
}
