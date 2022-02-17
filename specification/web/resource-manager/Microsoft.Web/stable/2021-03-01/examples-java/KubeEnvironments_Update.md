Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.models.KubeEnvironmentPatchResource;

/** Samples for KubeEnvironments Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/KubeEnvironments_Update.json
     */
    /**
     * Sample code: Update kube environments.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateKubeEnvironments(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getKubeEnvironments()
            .updateWithResponse(
                "examplerg", "testkubeenv", new KubeEnvironmentPatchResource().withStaticIp("1.2.3.4"), Context.NONE);
    }
}
```
