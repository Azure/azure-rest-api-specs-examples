Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for KubeEnvironments GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/KubeEnvironments_Get.json
     */
    /**
     * Sample code: Get kube environments by name.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getKubeEnvironmentsByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getKubeEnvironments()
            .getByResourceGroupWithResponse("examplerg", "jlaw-demo1", Context.NONE);
    }
}
```
