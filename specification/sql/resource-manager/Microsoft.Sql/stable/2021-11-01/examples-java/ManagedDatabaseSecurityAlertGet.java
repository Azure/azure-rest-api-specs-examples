
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;

/** Samples for ManagedDatabaseSecurityAlertPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseSecurityAlertGet.json
     */
    /**
     * Sample code: Get a database's threat detection policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getADatabaseSThreatDetectionPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSecurityAlertPolicies().getWithResponse(
            "securityalert-6852", "securityalert-2080", "testdb", SecurityAlertPolicyName.DEFAULT, Context.NONE);
    }
}
