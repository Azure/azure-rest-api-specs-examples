
/**
 * Samples for Fleets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01-preview/Fleets_Get.json
     */
    /**
     * Sample code: Gets a Fleet resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void
        getsAFleetResource(com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleets().getByResourceGroupWithResponse("rgfleets", "fleet1", com.azure.core.util.Context.NONE);
    }
}
