
/**
 * Samples for Profiles MigrationCommit.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Profiles_MigrationCommit.json
     */
    /**
     * Sample code: Profiles_MigrationCommit.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void profilesMigrationCommit(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getProfiles().migrationCommit("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}
