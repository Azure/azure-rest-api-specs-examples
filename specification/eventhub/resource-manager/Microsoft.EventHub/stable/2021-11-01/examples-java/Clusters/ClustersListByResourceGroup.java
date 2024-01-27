
import com.azure.core.util.Context;

/** Samples for Clusters ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/Clusters/
     * ClustersListByResourceGroup.json
     */
    /**
     * Sample code: ClustersListByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void clustersListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getClusters().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
