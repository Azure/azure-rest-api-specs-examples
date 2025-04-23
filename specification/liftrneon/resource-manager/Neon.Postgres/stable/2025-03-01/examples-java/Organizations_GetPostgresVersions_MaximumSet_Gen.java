
import com.azure.resourcemanager.neonpostgres.models.PgVersion;

/**
 * Samples for Organizations GetPostgresVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Organizations_GetPostgresVersions_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_GetPostgresVersions_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void
        organizationsGetPostgresVersionsMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.organizations().getPostgresVersionsWithResponse("rgneon", new PgVersion().withVersion(7),
            com.azure.core.util.Context.NONE);
    }
}
