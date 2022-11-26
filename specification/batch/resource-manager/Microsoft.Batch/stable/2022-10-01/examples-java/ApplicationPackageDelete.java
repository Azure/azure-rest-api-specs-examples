import com.azure.core.util.Context;

/** Samples for ApplicationPackage Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/ApplicationPackageDelete.json
     */
    /**
     * Sample code: ApplicationPackageDelete.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void applicationPackageDelete(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .applicationPackages()
            .deleteWithResponse("default-azurebatch-japaneast", "sampleacct", "app1", "1", Context.NONE);
    }
}
