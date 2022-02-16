Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.resources.models.ExportTemplateRequest;
import java.util.Arrays;

/** Samples for ResourceGroups ExportTemplate. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/ExportResourceGroup.json
     */
    /**
     * Sample code: Export a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void exportAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .serviceClient()
            .getResourceGroups()
            .exportTemplate(
                "my-resource-group",
                new ExportTemplateRequest()
                    .withResources(Arrays.asList("*"))
                    .withOptions("IncludeParameterDefaultValue,IncludeComments"),
                Context.NONE);
    }
}
```
