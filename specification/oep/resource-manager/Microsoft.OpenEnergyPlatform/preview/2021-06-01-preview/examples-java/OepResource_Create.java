/** Samples for EnergyServices Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/OepResource_Create.json
     */
    /**
     * Sample code: OepResource_Create.
     *
     * @param manager Entry point to OepManager.
     */
    public static void oepResourceCreate(com.azure.resourcemanager.oep.OepManager manager) {
        manager
            .energyServices()
            .define("DummyResourceName")
            .withRegion((String) null)
            .withExistingResourceGroup("DummyResourceGroupName")
            .create();
    }
}
