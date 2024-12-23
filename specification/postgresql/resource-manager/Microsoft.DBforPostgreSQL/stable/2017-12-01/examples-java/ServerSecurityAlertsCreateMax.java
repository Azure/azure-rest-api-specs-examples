
import com.azure.resourcemanager.postgresql.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.postgresql.models.ServerSecurityAlertPolicy;
import com.azure.resourcemanager.postgresql.models.ServerSecurityAlertPolicyState;
import java.util.Arrays;

/**
 * Samples for ServerSecurityAlertPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/
     * ServerSecurityAlertsCreateMax.json
     */
    /**
     * Sample code: Update a server's threat detection policy with all parameters.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void updateAServerSThreatDetectionPolicyWithAllParameters(
        com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        ServerSecurityAlertPolicy resource = manager.serverSecurityAlertPolicies().getWithResponse("securityalert-4799",
            "securityalert-6440", SecurityAlertPolicyName.DEFAULT, com.azure.core.util.Context.NONE).getValue();
        resource.update().withState(ServerSecurityAlertPolicyState.ENABLED)
            .withDisabledAlerts(Arrays.asList("Access_Anomaly", "Usage_Anomaly"))
            .withEmailAddresses(Arrays.asList("testSecurityAlert@microsoft.com")).withEmailAccountAdmins(true)
            .withStorageEndpoint("https://mystorage.blob.core.windows.net")
            .withStorageAccountAccessKey(
                "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==")
            .withRetentionDays(5).apply();
    }
}
