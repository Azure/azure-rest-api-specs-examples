
import com.azure.resourcemanager.synapse.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.synapse.models.SecurityAlertPolicyState;
import com.azure.resourcemanager.synapse.models.SqlPoolSecurityAlertPolicy;
import java.util.Arrays;

/**
 * Samples for SqlPoolSecurityAlertPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * CreateOrUpdateSqlPoolSecurityAlertWithAllParameters.json
     */
    /**
     * Sample code: Update a Sql pool's threat detection policy with all parameters.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void updateASqlPoolSThreatDetectionPolicyWithAllParameters(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        SqlPoolSecurityAlertPolicy resource
            = manager.sqlPoolSecurityAlertPolicies().getWithResponse("securityalert-4799", "securityalert-6440",
                "testdb", SecurityAlertPolicyName.DEFAULT, com.azure.core.util.Context.NONE).getValue();
        resource.update().withState(SecurityAlertPolicyState.ENABLED)
            .withDisabledAlerts(Arrays.asList("Sql_Injection", "Usage_Anomaly"))
            .withEmailAddresses(Arrays.asList("test@microsoft.com", "user@microsoft.com")).withEmailAccountAdmins(true)
            .withStorageEndpoint("https://mystorage.blob.core.windows.net")
            .withStorageAccountAccessKey(
                "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==")
            .withRetentionDays(6).apply();
    }
}
