Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Replications Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/ReplicationGet.json
     */
    /**
     * Sample code: ReplicationGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void replicationGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getReplications()
            .getWithResponse("myResourceGroup", "myRegistry", "myReplication", Context.NONE);
    }
}
```
