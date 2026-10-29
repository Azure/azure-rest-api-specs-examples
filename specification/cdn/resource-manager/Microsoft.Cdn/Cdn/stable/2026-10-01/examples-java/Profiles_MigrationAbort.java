
/**
 * Samples for Profiles MigrationAbort.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Profiles_MigrationAbort.json
     */
    /**
     * Sample code: Profiles_CdnMigrationCommit.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void profilesCdnMigrationCommit(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getProfiles().migrationAbort("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}
