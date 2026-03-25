
/**
 * Samples for FleetManagedNamespaces ListByFleet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01-preview/FleetManagedNamespaces_ListByFleet.json
     */
    /**
     * Sample code: FleetManagedNamespaces_ListByFleet.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void fleetManagedNamespacesListByFleet(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetManagedNamespaces().listByFleet("rgfleets", "fleet1", com.azure.core.util.Context.NONE);
    }
}
