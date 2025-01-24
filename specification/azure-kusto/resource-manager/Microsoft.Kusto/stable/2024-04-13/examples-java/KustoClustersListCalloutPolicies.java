
/**
 * Samples for Clusters ListCalloutPolicies.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/
     * KustoClustersListCalloutPolicies.json
     */
    /**
     * Sample code: KustoClusterListCalloutPolicies.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterListCalloutPolicies(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().listCalloutPolicies("kustorptest", "kustoCluster", com.azure.core.util.Context.NONE);
    }
}
