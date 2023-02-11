/** Samples for WorkspaceManagedIdentitySqlControlSettings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetManagedIdentitySqlControlSettings.json
     */
    /**
     * Sample code: Get managed identity sql control settings.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getManagedIdentitySqlControlSettings(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedIdentitySqlControlSettings()
            .getWithResponse("resourceGroup1", "workspace1", com.azure.core.util.Context.NONE);
    }
}
