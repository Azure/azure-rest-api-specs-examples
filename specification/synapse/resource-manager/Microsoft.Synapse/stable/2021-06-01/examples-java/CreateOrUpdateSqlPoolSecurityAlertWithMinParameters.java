import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.synapse.models.SecurityAlertPolicyState;
import com.azure.resourcemanager.synapse.models.SqlPoolSecurityAlertPolicy;

/** Samples for SqlPoolSecurityAlertPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolSecurityAlertWithMinParameters.json
     */
    /**
     * Sample code: Update a Sql pool's threat detection policy with minimal parameters.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void updateASqlPoolSThreatDetectionPolicyWithMinimalParameters(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        SqlPoolSecurityAlertPolicy resource =
            manager
                .sqlPoolSecurityAlertPolicies()
                .getWithResponse(
                    "securityalert-4799", "securityalert-6440", "testdb", SecurityAlertPolicyName.DEFAULT, Context.NONE)
                .getValue();
        resource.update().withState(SecurityAlertPolicyState.ENABLED).apply();
    }
}
