
import com.azure.resourcemanager.kusto.models.CalloutPolicyToRemove;

/**
 * Samples for Clusters RemoveCalloutPolicy.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/
     * KustoClusterRemoveCalloutPolicy.json
     */
    /**
     * Sample code: KustoClusterDropCalloutPolicy.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterDropCalloutPolicy(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().removeCalloutPolicy("kustorptest", "kustoCluster",
            new CalloutPolicyToRemove().withCalloutId("*_kusto"), com.azure.core.util.Context.NONE);
    }
}
