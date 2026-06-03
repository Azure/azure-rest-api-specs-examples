
/**
 * Samples for FleetManagedNamespaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/FleetManagedNamespaces_Delete.json
     */
    /**
     * Sample code: FleetManagedNamespaces_Delete.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void fleetManagedNamespacesDelete(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetManagedNamespaces().delete("rgfleets", "fleet1", "namespace1", null,
            com.azure.core.util.Context.NONE);
    }
}
