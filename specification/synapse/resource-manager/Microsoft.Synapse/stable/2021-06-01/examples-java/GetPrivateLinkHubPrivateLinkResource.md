Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateLinkHubPrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetPrivateLinkHubPrivateLinkResource.json
     */
    /**
     * Sample code: Get private link resources for private link hub.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getPrivateLinkResourcesForPrivateLinkHub(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .privateLinkHubPrivateLinkResources()
            .getWithResponse("ExampleResourceGroup", "ExamplePrivateLinkHub", "sql", Context.NONE);
    }
}
```
