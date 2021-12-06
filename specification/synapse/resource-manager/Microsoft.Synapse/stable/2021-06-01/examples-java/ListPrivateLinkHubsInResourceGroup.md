Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateLinkHubs ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListPrivateLinkHubsInResourceGroup.json
     */
    /**
     * Sample code: List privateLinkHubs in resource group.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listPrivateLinkHubsInResourceGroup(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.privateLinkHubs().listByResourceGroup("resourceGroup1", Context.NONE);
    }
}
```
