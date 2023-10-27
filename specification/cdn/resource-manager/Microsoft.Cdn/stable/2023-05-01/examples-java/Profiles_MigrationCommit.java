/** Samples for Profiles MigrationCommit. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Profiles_MigrationCommit.json
     */
    /**
     * Sample code: Profiles_MigrationCommit.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profilesMigrationCommit(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getProfiles()
            .migrationCommit("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}
