
import com.azure.resourcemanager.sql.fluent.models.ManagedServerSecurityAlertPolicyInner;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyState;
import java.util.Arrays;

/**
 * Samples for ManagedServerSecurityAlertPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedServerSecurityAlertCreateMax.json
     */
    /**
     * Sample code: Update a managed server's threat detection policy with all parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAManagedServerSThreatDetectionPolicyWithAllParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedServerSecurityAlertPolicies().createOrUpdate("securityalert-4799",
            "securityalert-6440", SecurityAlertPolicyName.DEFAULT,
            new ManagedServerSecurityAlertPolicyInner().withState(SecurityAlertPolicyState.ENABLED)
                .withDisabledAlerts(Arrays.asList("Access_Anomaly", "Usage_Anomaly"))
                .withEmailAddresses(Arrays.asList("testSecurityAlert@microsoft.com")).withEmailAccountAdmins(true)
                .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                .withStorageAccountAccessKey("fakeTokenPlaceholder").withRetentionDays(5),
            com.azure.core.util.Context.NONE);
    }
}
