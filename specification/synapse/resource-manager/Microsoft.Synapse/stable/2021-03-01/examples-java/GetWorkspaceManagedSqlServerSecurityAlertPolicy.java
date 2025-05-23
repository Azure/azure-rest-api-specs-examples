
import com.azure.resourcemanager.synapse.models.SecurityAlertPolicyNameAutoGenerated;

/**
 * Samples for WorkspaceManagedSqlServerSecurityAlertPolicy Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * GetWorkspaceManagedSqlServerSecurityAlertPolicy.json
     */
    /**
     * Sample code: Get workspace managed sql Server's security alert policy.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void
        getWorkspaceManagedSqlServerSSecurityAlertPolicy(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaceManagedSqlServerSecurityAlertPolicies().getWithResponse("wsg-7398", "testWorkspace",
            SecurityAlertPolicyNameAutoGenerated.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
