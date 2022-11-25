import com.azure.core.util.Context;

/** Samples for Pool StopResize. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolStopResize.json
     */
    /**
     * Sample code: StopPoolResize.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void stopPoolResize(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().stopResizeWithResponse("default-azurebatch-japaneast", "sampleacct", "testpool", Context.NONE);
    }
}
