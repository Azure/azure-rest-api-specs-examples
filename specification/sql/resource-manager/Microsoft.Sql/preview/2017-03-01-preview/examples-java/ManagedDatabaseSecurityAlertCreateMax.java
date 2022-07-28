import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;

/** Samples for ManagedDatabaseSecurityAlertPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ManagedDatabaseSecurityAlertCreateMax.json
     */
    /**
     * Sample code: Update a database's threat detection policy with all parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateADatabaseSThreatDetectionPolicyWithAllParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedDatabaseSecurityAlertPolicies()
            .createOrUpdateWithResponse(
                "securityalert-4799",
                "securityalert-6440",
                "testdb",
                SecurityAlertPolicyName.DEFAULT,
                null,
                Context.NONE);
    }
}
