
import com.azure.resourcemanager.network.models.MigrationType;
import com.azure.resourcemanager.network.models.NetworkVirtualAppliancePrepareMigrationProperties;
import com.azure.resourcemanager.network.models.NetworkVirtualAppliancePrepareMigrationRequest;

/**
 * Samples for NetworkVirtualAppliances PrepareMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkVirtualAppliancePrepareMigration.json
     */
    /**
     * Sample code: Prepare migration of a NetworkVirtualAppliance to the new ILB architecture.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void prepareMigrationOfANetworkVirtualApplianceToTheNewILBArchitecture(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().prepareMigration("rg1", "nva",
            new NetworkVirtualAppliancePrepareMigrationRequest()
                .withProperties(new NetworkVirtualAppliancePrepareMigrationProperties()
                    .withMigrationType(MigrationType.MIGRATE_TO_NEW_ILBARCHITECTURE)),
            com.azure.core.util.Context.NONE);
    }
}
