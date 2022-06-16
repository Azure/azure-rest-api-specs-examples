import com.azure.core.util.Context;

/** Samples for WorkspaceManagedSqlServerEncryptionProtector List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlServerEncryptionProtectorList.json
     */
    /**
     * Sample code: Get workspace managed sql Server's encryption protectors.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getWorkspaceManagedSqlServerSEncryptionProtectors(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaceManagedSqlServerEncryptionProtectors().list("wsg-7398", "testWorkspace", Context.NONE);
    }
}
