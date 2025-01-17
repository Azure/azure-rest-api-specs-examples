
import com.azure.resourcemanager.synapse.fluent.models.ManagedIdentitySqlControlSettingsModelInner;
import com.azure.resourcemanager.synapse.models.ManagedIdentitySqlControlSettingsModelPropertiesGrantSqlControlToManagedIdentity;
import com.azure.resourcemanager.synapse.models.ManagedIdentitySqlControlSettingsModelPropertiesGrantSqlControlToManagedIdentityDesiredState;

/**
 * Samples for WorkspaceManagedIdentitySqlControlSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * CreateOrUpdateManagedIdentitySqlControlSettings.json
     */
    /**
     * Sample code: Create or update managed identity sql control settings.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void
        createOrUpdateManagedIdentitySqlControlSettings(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaceManagedIdentitySqlControlSettings().createOrUpdate("resourceGroup1", "workspace1",
            new ManagedIdentitySqlControlSettingsModelInner().withGrantSqlControlToManagedIdentity(
                new ManagedIdentitySqlControlSettingsModelPropertiesGrantSqlControlToManagedIdentity().withDesiredState(
                    ManagedIdentitySqlControlSettingsModelPropertiesGrantSqlControlToManagedIdentityDesiredState.ENABLED)),
            com.azure.core.util.Context.NONE);
    }
}
