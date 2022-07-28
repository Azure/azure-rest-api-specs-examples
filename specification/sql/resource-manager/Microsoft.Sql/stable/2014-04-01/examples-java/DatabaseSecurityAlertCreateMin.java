import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DatabaseSecurityAlertPolicyInner;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyState;

/** Samples for DatabaseThreatDetectionPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DatabaseSecurityAlertCreateMin.json
     */
    /**
     * Sample code: Create database security alert policy min.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createDatabaseSecurityAlertPolicyMin(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabaseThreatDetectionPolicies()
            .createOrUpdateWithResponse(
                "securityalert-4799",
                "securityalert-6440",
                "testdb",
                SecurityAlertPolicyName.DEFAULT,
                new DatabaseSecurityAlertPolicyInner()
                    .withState(SecurityAlertPolicyState.ENABLED)
                    .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                    .withStorageAccountAccessKey(
                        "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
                Context.NONE);
    }
}
