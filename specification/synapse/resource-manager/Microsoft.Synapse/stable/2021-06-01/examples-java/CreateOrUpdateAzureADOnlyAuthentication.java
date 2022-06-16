import com.azure.resourcemanager.synapse.models.AzureADOnlyAuthenticationName;

/** Samples for AzureADOnlyAuthentications Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateAzureADOnlyAuthentication.json
     */
    /**
     * Sample code: Create or Update Azure Active Directory Only Authentication property.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateAzureActiveDirectoryOnlyAuthenticationProperty(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .azureADOnlyAuthentications()
            .define(AzureADOnlyAuthenticationName.DEFAULT)
            .withExistingWorkspace("workspace-6852", "workspace-2080")
            .withAzureADOnlyAuthentication(true)
            .create();
    }
}
