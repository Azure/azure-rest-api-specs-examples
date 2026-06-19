
import com.azure.resourcemanager.managednetworkfabric.models.DiscardCommitBatchRequest;

/**
 * Samples for NetworkFabrics DiscardCommitBatch.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_DiscardCommitBatch.json
     */
    /**
     * Sample code: NetworkFabrics_DiscardCommitBatch.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsDiscardCommitBatch(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().discardCommitBatch("example-rg", "example-fabric",
            new DiscardCommitBatchRequest().withCommitBatchId("batchId1"), com.azure.core.util.Context.NONE);
    }
}
