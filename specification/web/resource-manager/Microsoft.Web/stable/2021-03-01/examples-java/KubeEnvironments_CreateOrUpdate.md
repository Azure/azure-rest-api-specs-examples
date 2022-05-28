Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.fluent.models.KubeEnvironmentInner;

/** Samples for KubeEnvironments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/KubeEnvironments_CreateOrUpdate.json
     */
    /**
     * Sample code: Create kube environments.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createKubeEnvironments(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getKubeEnvironments()
            .createOrUpdate(
                "examplerg",
                "testkubeenv",
                new KubeEnvironmentInner().withLocation("East US").withStaticIp("1.2.3.4"),
                Context.NONE);
    }
}
```
