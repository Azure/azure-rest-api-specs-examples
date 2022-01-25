Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.fluent.models.ManagedIdentitySqlControlSettingsModelInner;
import com.azure.resourcemanager.synapse.models.ManagedIdentitySqlControlSettingsModelPropertiesGrantSqlControlToManagedIdentity;
import com.azure.resourcemanager.synapse.models.ManagedIdentitySqlControlSettingsModelPropertiesGrantSqlControlToManagedIdentityDesiredState;

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
                            .withDesiredState(
                                ManagedIdentitySqlControlSettingsModelPropertiesGrantSqlControlToManagedIdentityDesiredState
                                    .ENABLED)),
                Context.NONE);
    }
}
```
