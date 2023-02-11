/** Samples for BigDataPools Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/DeleteBigDataPool.json
     */
    /**
     * Sample code: Delete a Big Data pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deleteABigDataPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .bigDataPools()
            .delete("ExampleResourceGroup", "ExampleWorkspace", "ExamplePool", com.azure.core.util.Context.NONE);
    }
}
