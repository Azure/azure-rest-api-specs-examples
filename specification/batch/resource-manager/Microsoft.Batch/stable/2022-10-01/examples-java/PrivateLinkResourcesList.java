import com.azure.core.util.Context;

/** Samples for PrivateLinkResource ListByBatchAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PrivateLinkResourcesList.json
     */
    /**
     * Sample code: ListPrivateLinkResource.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void listPrivateLinkResource(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .privateLinkResources()
            .listByBatchAccount("default-azurebatch-japaneast", "sampleacct", null, Context.NONE);
    }
}
