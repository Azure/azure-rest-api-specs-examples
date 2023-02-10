import com.azure.resourcemanager.synapse.models.EncryptionProtector;
import com.azure.resourcemanager.synapse.models.EncryptionProtectorName;
import com.azure.resourcemanager.synapse.models.ServerKeyType;

/** Samples for WorkspaceManagedSqlServerEncryptionProtector CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateServiceManaged.json
     */
    /**
     * Sample code: Update the encryption protector to service managed.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void updateTheEncryptionProtectorToServiceManaged(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        EncryptionProtector resource =
            manager
                .workspaceManagedSqlServerEncryptionProtectors()
                .getWithResponse(
                    "wsg-7398", "testWorkspace", EncryptionProtectorName.CURRENT, com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withServerKeyName("ServiceManaged").withServerKeyType(ServerKeyType.SERVICE_MANAGED).apply();
    }
}
