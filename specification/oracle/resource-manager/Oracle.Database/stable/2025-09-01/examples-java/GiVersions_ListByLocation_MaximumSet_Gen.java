
import com.azure.resourcemanager.oracledatabase.models.SystemShapes;

/**
 * Samples for GiVersions ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/GiVersions_ListByLocation_MaximumSet_Gen.json
     */
    /**
     * Sample code: GiVersions_ListByLocation_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        giVersionsListByLocationMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.giVersions().listByLocation("eastus", SystemShapes.EXADATA_X9M, "hpzuyaemum", null,
            com.azure.core.util.Context.NONE);
    }
}
