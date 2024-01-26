
import com.azure.core.util.Context;

/** Samples for Clusters GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/Clusters/ClusterGet.json
     */
    /**
     * Sample code: ClusterGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void clusterGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getClusters().getByResourceGroupWithResponse("myResourceGroup",
            "testCluster", Context.NONE);
    }
}
