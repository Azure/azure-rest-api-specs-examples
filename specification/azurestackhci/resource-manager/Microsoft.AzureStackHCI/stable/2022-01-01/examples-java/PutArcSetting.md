Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-azurestackhci_1.0.0-beta.2/sdk/azurestackhci/azure-resourcemanager-azurestackhci/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for ArcSettings Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-01-01/examples/PutArcSetting.json
     */
    /**
     * Sample code: Create ArcSetting.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void createArcSetting(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.arcSettings().define("default").withExistingCluster("test-rg", "myCluster").create();
    }
}
```
