import com.azure.core.util.Context;

/** Samples for EnergyServices GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/OepResource_Get.json
     */
    /**
     * Sample code: OepResource_Get.
     *
     * @param manager Entry point to OepManager.
     */
    public static void oepResourceGet(com.azure.resourcemanager.oep.OepManager manager) {
        manager
            .energyServices()
            .getByResourceGroupWithResponse("DummyResourceGroupName", "DummyResourceName", Context.NONE);
    }
}
