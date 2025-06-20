
import com.azure.resourcemanager.oracledatabase.models.ShapeFamily;

/**
 * Samples for GiMinorVersions ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/GiMinorVersions_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: GiMinorVersions_ListByParent_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        giMinorVersionsListByParentMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.giMinorVersions().listByParent("eastus", "giVersionName",
            ShapeFamily.fromString("rtfcosvtlpeeqoicsjqggtgc"), null, com.azure.core.util.Context.NONE);
    }
}
