Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WebApps GetAppSettingKeyVaultReferenceSlot. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetKeyVaultReferencesForAppSettingSlot.json
     */
    /**
     * Sample code: Get Azure Key Vault slot app setting reference.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAzureKeyVaultSlotAppSettingReference(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWebApps()
            .getAppSettingKeyVaultReferenceSlotWithResponse("testrg123", "testc6282", "setting", "stage", Context.NONE);
    }
}
```
