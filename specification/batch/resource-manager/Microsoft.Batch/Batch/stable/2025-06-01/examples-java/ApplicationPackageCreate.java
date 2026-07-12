
/**
 * Samples for ApplicationPackage Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/ApplicationPackageCreate.json
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
