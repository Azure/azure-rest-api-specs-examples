
import com.azure.core.util.Context;

/** Samples for ServerSecurityAlertPolicies ListByServer. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerSecurityAlertsListByServer.json
     */
    /**
     * Sample code: List the server's threat detection policies.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTheServerSThreatDetectionPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerSecurityAlertPolicies().listByServer("securityalert-4799",
            "securityalert-6440", Context.NONE);
    }
}
