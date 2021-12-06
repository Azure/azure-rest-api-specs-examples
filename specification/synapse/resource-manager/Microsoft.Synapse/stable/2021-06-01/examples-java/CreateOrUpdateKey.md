Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
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
```
