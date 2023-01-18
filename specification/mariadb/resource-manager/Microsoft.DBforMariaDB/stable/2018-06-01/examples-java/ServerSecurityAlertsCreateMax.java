import com.azure.resourcemanager.mariadb.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.mariadb.models.ServerSecurityAlertPolicy;
import com.azure.resourcemanager.mariadb.models.ServerSecurityAlertPolicyState;
import java.util.Arrays;

/** Samples for ServerSecurityAlertPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerSecurityAlertsCreateMax.json
     */
    /**
     * Sample code: Update a server's threat detection policy with all parameters.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void updateAServerSThreatDetectionPolicyWithAllParameters(
        com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        ServerSecurityAlertPolicy resource =
            manager
                .serverSecurityAlertPolicies()
                .getWithResponse(
                    "securityalert-4799",
                    "securityalert-6440",
                    SecurityAlertPolicyName.DEFAULT,
                    com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withState(ServerSecurityAlertPolicyState.ENABLED)
            .withDisabledAlerts(Arrays.asList("Access_Anomaly", "Usage_Anomaly"))
            .withEmailAddresses(Arrays.asList("testSecurityAlert@microsoft.com"))
            .withEmailAccountAdmins(true)
            .withStorageEndpoint("https://mystorage.blob.core.windows.net")
            .withStorageAccountAccessKey(
                "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==")
            .withRetentionDays(5)
            .apply();
    }
}
