
import com.azure.core.util.Context;

/** Samples for Clusters ListNamespaces. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/Clusters/
     * ListNamespacesInClusterGet.json
     */
    /**
     * Sample code: ListNamespacesInCluster.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNamespacesInCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getClusters().listNamespacesWithResponse("myResourceGroup",
            "testCluster", Context.NONE);
    }
}
