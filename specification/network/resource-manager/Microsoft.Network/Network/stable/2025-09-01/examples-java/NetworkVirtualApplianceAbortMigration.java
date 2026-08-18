
/**
 * Samples for NetworkVirtualAppliances AbortMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkVirtualApplianceAbortMigration.json
     */
    /**
     * Sample code: Abort migration of a NetworkVirtualAppliance to the new ILB architecture.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void abortMigrationOfANetworkVirtualApplianceToTheNewILBArchitecture(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().abortMigration("rg1", "nva",
            com.azure.core.util.Context.NONE);
    }
}
