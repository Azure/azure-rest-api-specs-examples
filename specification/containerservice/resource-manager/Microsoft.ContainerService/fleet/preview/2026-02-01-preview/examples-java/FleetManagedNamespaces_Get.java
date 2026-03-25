
/**
 * Samples for FleetManagedNamespaces Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01-preview/FleetManagedNamespaces_Get.json
     */
    /**
     * Sample code: FleetManagedNamespaces_Get.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void fleetManagedNamespacesGet(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetManagedNamespaces().getWithResponse("rgfleets", "fleet1", "namespace1",
            com.azure.core.util.Context.NONE);
    }
}
