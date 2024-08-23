
import com.azure.resourcemanager.sql.fluent.models.ManagedDatabaseSecurityAlertPolicyInner;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyState;

/**
 * Samples for ManagedDatabaseSecurityAlertPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseSecurityAlertCreateMin
     * .json
     */
    /**
     * Sample code: Update a database's threat detection policy with minimal parameters.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateADatabaseSThreatDetectionPolicyWithMinimalParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSecurityAlertPolicies()
            .createOrUpdateWithResponse("securityalert-4799", "securityalert-6440", "testdb",
                SecurityAlertPolicyName.DEFAULT,
                new ManagedDatabaseSecurityAlertPolicyInner().withState(SecurityAlertPolicyState.ENABLED),
                com.azure.core.util.Context.NONE);
    }
}
