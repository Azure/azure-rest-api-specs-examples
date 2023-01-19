/** Samples for Operations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetFailedOperationResult.json
     */
    /**
     * Sample code: GetFailedOperationResult.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void getFailedOperationResult(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .operations()
            .getWithResponse(
                "westus2",
                "https://management.azure.com/",
                "d030bb3f-7d53-11e9-8e09-9a86872085ff",
                com.azure.core.util.Context.NONE);
    }
}
