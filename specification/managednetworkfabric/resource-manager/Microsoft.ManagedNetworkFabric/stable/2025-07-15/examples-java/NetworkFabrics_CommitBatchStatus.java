
import com.azure.resourcemanager.managednetworkfabric.models.CommitBatchStatusRequest;

/**
 * Samples for NetworkFabrics CommitBatchStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_CommitBatchStatus.json
     */
    /**
     * Sample code: NetworkFabrics_CommitBatchStatus_MaximumSet.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsCommitBatchStatusMaximumSet(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().commitBatchStatus("example-rg", "example-fabric",
            new CommitBatchStatusRequest().withCommitBatchId("batch-id"), com.azure.core.util.Context.NONE);
    }
}
