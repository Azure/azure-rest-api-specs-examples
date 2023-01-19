/** Samples for PeeringServices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/CreatePeeringService.json
     */
    /**
     * Sample code: Create a peering service.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void createAPeeringService(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager
            .peeringServices()
            .define("peeringServiceName")
            .withRegion("eastus")
            .withExistingResourceGroup("rgName")
            .withPeeringServiceLocation("state1")
            .withPeeringServiceProvider("serviceProvider1")
            .withProviderPrimaryPeeringLocation("peeringLocation1")
            .withProviderBackupPeeringLocation("peeringLocation2")
            .create();
    }
}
