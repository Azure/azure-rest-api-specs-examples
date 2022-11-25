import com.azure.core.util.Context;

/** Samples for Application List. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/ApplicationList.json
     */
    /**
     * Sample code: ApplicationList.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void applicationList(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.applications().list("default-azurebatch-japaneast", "sampleacct", null, Context.NONE);
    }
}
