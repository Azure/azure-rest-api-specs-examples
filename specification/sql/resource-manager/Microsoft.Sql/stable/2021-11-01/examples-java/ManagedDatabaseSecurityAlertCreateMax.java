
import com.azure.resourcemanager.sql.fluent.models.ManagedDatabaseSecurityAlertPolicyInner;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyState;
import java.util.Arrays;

/**
 * Samples for ManagedDatabaseSecurityAlertPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseSecurityAlertCreateMax
     * .json
     */
    /**
     * Sample code: Update a database's threat detection policy with all parameters.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        updateADatabaseSThreatDetectionPolicyWithAllParameters(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSecurityAlertPolicies()
            .createOrUpdateWithResponse("securityalert-4799", "securityalert-6440", "testdb",
                SecurityAlertPolicyName.DEFAULT,
                new ManagedDatabaseSecurityAlertPolicyInner().withState(SecurityAlertPolicyState.ENABLED)
                    .withDisabledAlerts(Arrays.asList("Sql_Injection", "Usage_Anomaly"))
                    .withEmailAddresses(Arrays.asList("test@contoso.com", "user@contoso.com"))
                    .withEmailAccountAdmins(true).withStorageEndpoint("https://mystorage.blob.core.windows.net")
                    .withStorageAccountAccessKey("fakeTokenPlaceholder").withRetentionDays(6),
                com.azure.core.util.Context.NONE);
    }
}
