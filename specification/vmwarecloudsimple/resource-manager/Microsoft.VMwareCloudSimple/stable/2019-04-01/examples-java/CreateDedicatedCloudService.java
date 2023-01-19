/** Samples for DedicatedCloudServices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/CreateDedicatedCloudService.json
     */
    /**
     * Sample code: CreateDedicatedCloudService.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void createDedicatedCloudService(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .dedicatedCloudServices()
            .define("myService")
            .withRegion("westus")
            .withExistingResourceGroup("myResourceGroup")
            .withGatewaySubnet("10.0.0.0")
            .create();
    }
}
