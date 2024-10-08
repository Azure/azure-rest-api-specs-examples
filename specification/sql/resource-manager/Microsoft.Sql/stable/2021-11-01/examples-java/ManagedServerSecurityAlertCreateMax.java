
import com.azure.resourcemanager.sql.fluent.models.ManagedServerSecurityAlertPolicyInner;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyNameAutoGenerated;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyState;
import java.util.Arrays;

/**
 * Samples for ManagedServerSecurityAlertPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedServerSecurityAlertCreateMax.
     * json
     */
    /**
     * Sample code: Update a managed server's threat detection policy with all parameters.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAManagedServerSThreatDetectionPolicyWithAllParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedServerSecurityAlertPolicies().createOrUpdate(
            "securityalert-4799", "securityalert-6440", SecurityAlertPolicyNameAutoGenerated.DEFAULT,
            new ManagedServerSecurityAlertPolicyInner().withState(SecurityAlertPolicyState.ENABLED)
                .withDisabledAlerts(Arrays.asList("Access_Anomaly", "Usage_Anomaly"))
                .withEmailAddresses(Arrays.asList("testSecurityAlert@microsoft.com")).withEmailAccountAdmins(true)
                .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                .withStorageAccountAccessKey("fakeTokenPlaceholder").withRetentionDays(5),
            com.azure.core.util.Context.NONE);
    }
}
