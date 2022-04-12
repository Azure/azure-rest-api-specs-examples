Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AgentPools UpgradeNodeImageVersion. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-02-01/examples/AgentPoolsUpgradeNodeImageVersion.json
     */
    /**
     * Sample code: Upgrade Agent Pool Node Image Version.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void upgradeAgentPoolNodeImageVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getAgentPools()
            .upgradeNodeImageVersion("rg1", "clustername1", "agentpool1", Context.NONE);
    }
}
```
