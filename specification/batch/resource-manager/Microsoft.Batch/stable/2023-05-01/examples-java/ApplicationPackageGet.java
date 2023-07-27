/** Samples for ApplicationPackage Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/ApplicationPackageGet.json
     */
    /**
     * Sample code: ApplicationPackageGet.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void applicationPackageGet(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .applicationPackages()
            .getWithResponse(
                "default-azurebatch-japaneast", "sampleacct", "app1", "1", com.azure.core.util.Context.NONE);
    }
}
