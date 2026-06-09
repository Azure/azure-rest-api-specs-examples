
/**
 * Samples for MajorVersionUpgradePrecheck List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/MajorVersionUpgradePrecheckListByServer.json
     */
    /**
     * Sample code: List all major version upgrade precheck validations for a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listAllMajorVersionUpgradePrecheckValidationsForAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.majorVersionUpgradePrechecks().list("exampleresourcegroup", "exampleserver",
            com.azure.core.util.Context.NONE);
    }
}
