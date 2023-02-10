import com.azure.resourcemanager.synapse.models.EncryptionProtectorName;

/** Samples for WorkspaceManagedSqlServerEncryptionProtector Revalidate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/WorkspaceManagedSqlServerEncryptionProtectorRevalidate.json
     */
    /**
     * Sample code: Revalidates the encryption protector.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void revalidatesTheEncryptionProtector(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerEncryptionProtectors()
            .revalidate("wsg-7398", "testWorkspace", EncryptionProtectorName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
