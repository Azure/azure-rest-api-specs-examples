
/**
 * Samples for PrivateLinkResource ListByBatchAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/PrivateLinkResourcesList.json
     */
    /**
     * Sample code: ListPrivateLinkResource.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void listPrivateLinkResource(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.privateLinkResources().listByBatchAccount("default-azurebatch-japaneast", "sampleacct", null,
            com.azure.core.util.Context.NONE);
    }
}
