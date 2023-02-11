import com.azure.resourcemanager.synapse.fluent.models.ManagedIdentitySqlControlSettingsModelInner;
import com.azure.resourcemanager.synapse.models.DesiredState;
import com.azure.resourcemanager.synapse.models.ManagedIdentitySqlControlSettingsModelPropertiesGrantSqlControlToManagedIdentity;

/** Samples for WorkspaceManagedIdentitySqlControlSettings CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateManagedIdentitySqlControlSettings.json
     */
    /**
     * Sample code: Create or update managed identity sql control settings.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateManagedIdentitySqlControlSettings(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedIdentitySqlControlSettings()
            .createOrUpdate(
                "resourceGroup1",
                "workspace1",
                new ManagedIdentitySqlControlSettingsModelInner()
                    .withGrantSqlControlToManagedIdentity(
                        new ManagedIdentitySqlControlSettingsModelPropertiesGrantSqlControlToManagedIdentity()
                            .withDesiredState(DesiredState.ENABLED)),
                com.azure.core.util.Context.NONE);
    }
}
