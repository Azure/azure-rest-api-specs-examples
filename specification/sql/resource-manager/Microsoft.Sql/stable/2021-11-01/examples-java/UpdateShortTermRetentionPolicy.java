
import com.azure.resourcemanager.sql.fluent.models.BackupShortTermRetentionPolicyInner;
import com.azure.resourcemanager.sql.models.DiffBackupIntervalInHours;
import com.azure.resourcemanager.sql.models.ShortTermRetentionPolicyName;

/**
 * Samples for BackupShortTermRetentionPolicies Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/UpdateShortTermRetentionPolicy.json
     */
    /**
     * Sample code: Update the short term retention policy for the database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        updateTheShortTermRetentionPolicyForTheDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getBackupShortTermRetentionPolicies().update("resourceGroup",
            "testsvr", "testdb", ShortTermRetentionPolicyName.DEFAULT, new BackupShortTermRetentionPolicyInner()
                .withRetentionDays(7).withDiffBackupIntervalInHours(DiffBackupIntervalInHours.TWO_FOUR),
            com.azure.core.util.Context.NONE);
    }
}
