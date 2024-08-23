
import com.azure.resourcemanager.sql.fluent.models.LongTermRetentionPolicyInner;
import com.azure.resourcemanager.sql.models.LongTermRetentionPolicyName;

/**
 * Samples for LongTermRetentionPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/LongTermRetentionPolicyCreateOrUpdate
     * .json
     */
    /**
     * Sample code: Create or update the long term retention policy for the database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateTheLongTermRetentionPolicyForTheDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getLongTermRetentionPolicies().createOrUpdate("resourceGroup",
            "testserver", "testDatabase", LongTermRetentionPolicyName.DEFAULT, new LongTermRetentionPolicyInner()
                .withWeeklyRetention("P1M").withMonthlyRetention("P1Y").withYearlyRetention("P5Y").withWeekOfYear(5),
            com.azure.core.util.Context.NONE);
    }
}
