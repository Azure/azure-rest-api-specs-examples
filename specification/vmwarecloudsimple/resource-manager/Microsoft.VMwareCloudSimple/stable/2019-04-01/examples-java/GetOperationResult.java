/** Samples for Operations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetOperationResult.json
     */
    /**
     * Sample code: GetOperationResult.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void getOperationResult(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .operations()
            .getWithResponse(
                "westus2",
                "https://management.azure.com/",
                "f8e1c8f1-7d52-11e9-8e07-9a86872085ff",
                com.azure.core.util.Context.NONE);
    }
}
