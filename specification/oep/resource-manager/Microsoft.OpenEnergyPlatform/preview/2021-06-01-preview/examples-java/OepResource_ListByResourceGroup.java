import com.azure.core.util.Context;

/** Samples for EnergyServices ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/OepResource_ListByResourceGroup.json
     */
    /**
     * Sample code: OepResource_ListByResourceGroup.
     *
     * @param manager Entry point to OepManager.
     */
    public static void oepResourceListByResourceGroup(com.azure.resourcemanager.oep.OepManager manager) {
        manager.energyServices().listByResourceGroup("DummyResourceGroupName", Context.NONE);
    }
}
