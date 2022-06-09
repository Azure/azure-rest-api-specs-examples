```java
import com.azure.core.util.Context;

/** Samples for WebApps GetAppSettingKeyVaultReference. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetKeyVaultReferencesForAppSetting.json
     */
    /**
     * Sample code: Get Azure Key Vault app setting reference.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAzureKeyVaultAppSettingReference(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWebApps()
            .getAppSettingKeyVaultReferenceWithResponse("testrg123", "testc6282", "setting", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
