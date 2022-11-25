import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/OperationsList.json
     */
    /**
     * Sample code: OperationsList.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void operationsList(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.operations().list(Context.NONE);
    }
}
