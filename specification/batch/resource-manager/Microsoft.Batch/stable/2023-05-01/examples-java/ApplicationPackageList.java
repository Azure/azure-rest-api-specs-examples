/** Samples for ApplicationPackage List. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/ApplicationPackageList.json
     */
    /**
     * Sample code: ApplicationPackageList.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void applicationPackageList(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .applicationPackages()
            .list("default-azurebatch-japaneast", "sampleacct", "app1", null, com.azure.core.util.Context.NONE);
    }
}
