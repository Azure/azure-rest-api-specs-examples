
import com.azure.resourcemanager.network.models.MigrationType;
import com.azure.resourcemanager.network.models.NetworkVirtualApplianceCommitMigrationProperties;
import com.azure.resourcemanager.network.models.NetworkVirtualApplianceCommitMigrationRequest;

/**
 * Samples for NetworkVirtualAppliances CommitMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkVirtualApplianceCommitMigration.json
     */
    /**
     * Sample code: Commit migration of a NetworkVirtualAppliance to the new ILB architecture.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void commitMigrationOfANetworkVirtualApplianceToTheNewILBArchitecture(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().commitMigration("rg1", "nva",
            new NetworkVirtualApplianceCommitMigrationRequest()
                .withProperties(new NetworkVirtualApplianceCommitMigrationProperties()
                    .withMigrationType(MigrationType.MIGRATE_TO_NEW_ILBARCHITECTURE)),
            com.azure.core.util.Context.NONE);
    }
}
