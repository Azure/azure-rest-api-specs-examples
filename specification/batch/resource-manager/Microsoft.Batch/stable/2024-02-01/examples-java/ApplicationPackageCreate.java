
/**
 * Samples for ApplicationPackage Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/ApplicationPackageCreate.json
     */
    /**
     * Sample code: ApplicationPackageCreate.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void applicationPackageCreate(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.applicationPackages().define("1")
            .withExistingApplication("default-azurebatch-japaneast", "sampleacct", "app1").create();
    }
}
