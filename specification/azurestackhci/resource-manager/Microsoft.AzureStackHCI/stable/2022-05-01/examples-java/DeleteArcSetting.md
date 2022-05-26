Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-azurestackhci_1.0.0-beta.3/sdk/azurestackhci/azure-resourcemanager-azurestackhci/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ArcSettings Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/DeleteArcSetting.json
     */
    /**
     * Sample code: Delete ArcSetting.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteArcSetting(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.arcSettings().delete("test-rg", "myCluster", "default", Context.NONE);
    }
}
```
