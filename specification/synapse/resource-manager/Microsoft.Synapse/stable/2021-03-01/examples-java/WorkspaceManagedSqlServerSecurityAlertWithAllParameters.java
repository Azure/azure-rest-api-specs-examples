
import com.azure.resourcemanager.synapse.models.SecurityAlertPolicyNameAutoGenerated;
import com.azure.resourcemanager.synapse.models.SecurityAlertPolicyState;
import com.azure.resourcemanager.synapse.models.ServerSecurityAlertPolicy;
import java.util.Arrays;

/**
 * Samples for WorkspaceManagedSqlServerSecurityAlertPolicy CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * WorkspaceManagedSqlServerSecurityAlertWithAllParameters.json
     */
    /**
     * Sample code: Update a workspace managed sql server's threat detection policy with all parameters.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void updateAWorkspaceManagedSqlServerSThreatDetectionPolicyWithAllParameters(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        ServerSecurityAlertPolicy resource
            = manager.workspaceManagedSqlServerSecurityAlertPolicies().getWithResponse("wsg-7398", "testWorkspace",
                SecurityAlertPolicyNameAutoGenerated.DEFAULT, com.azure.core.util.Context.NONE).getValue();
        resource.update().withState(SecurityAlertPolicyState.ENABLED)
            .withDisabledAlerts(Arrays.asList("Access_Anomaly", "Usage_Anomaly"))
            .withEmailAddresses(Arrays.asList("testSecurityAlert@microsoft.com")).withEmailAccountAdmins(true)
            .withStorageEndpoint("https://mystorage.blob.core.windows.net")
            .withStorageAccountAccessKey(
                "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==")
            .withRetentionDays(5).apply();
    }
}
