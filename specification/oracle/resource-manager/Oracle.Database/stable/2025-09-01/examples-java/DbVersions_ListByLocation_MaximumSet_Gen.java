
import com.azure.resourcemanager.oracledatabase.models.BaseDbSystemShapes;
import com.azure.resourcemanager.oracledatabase.models.ShapeFamilyType;
import com.azure.resourcemanager.oracledatabase.models.StorageManagementType;

/**
 * Samples for DbVersions ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/DbVersions_ListByLocation_MaximumSet_Gen.json
     */
    /**
     * Sample code: DbVersions_ListByLocation_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        dbVersionsListByLocationMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbVersions().listByLocation("eastus", BaseDbSystemShapes.VMSTANDARD_X86,
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Oracle.Database/dbSystems/dbsystem1",
            StorageManagementType.LVM, true, true, ShapeFamilyType.VIRTUAL_MACHINE, com.azure.core.util.Context.NONE);
    }
}
