import com.azure.core.util.Context;

/** Samples for Pool Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolDelete.json
     */
    /**
     * Sample code: DeletePool.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void deletePool(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().delete("default-azurebatch-japaneast", "sampleacct", "testpool", Context.NONE);
    }
}
