/** Samples for Keys CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateKey.json
     */
    /**
     * Sample code: Create or update a workspace key.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateAWorkspaceKey(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .keys()
            .define("somekey")
            .withExistingWorkspace("ExampleResourceGroup", "ExampleWorkspace")
            .withIsActiveCmk(true)
            .withKeyVaultUrl("https://vault.azure.net/keys/somesecret")
            .create();
    }
}
