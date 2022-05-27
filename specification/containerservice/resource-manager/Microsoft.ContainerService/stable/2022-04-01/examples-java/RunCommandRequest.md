Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.containerservice.models.RunCommandRequest;

/** Samples for ManagedClusters RunCommand. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-04-01/examples/RunCommandRequest.json
     */
    /**
     * Sample code: submitNewCommand.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void submitNewCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getManagedClusters()
            .runCommand(
                "rg1",
                "clustername1",
                new RunCommandRequest().withCommand("kubectl apply -f ns.yaml").withContext("").withClusterToken(""),
                Context.NONE);
    }
}
```
