Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.apimanagement.models.KeyVaultContractCreateProperties;
import java.util.Arrays;

/** Samples for NamedValue CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateNamedValueWithKeyVault.json
     */
    /**
     * Sample code: ApiManagementCreateNamedValueWithKeyVault.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateNamedValueWithKeyVault(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .namedValues()
            .define("testprop6")
            .withExistingService("rg1", "apimService1")
            .withTags(Arrays.asList("foo", "bar"))
            .withDisplayName("prop6namekv")
            .withKeyVault(
                new KeyVaultContractCreateProperties()
                    .withSecretIdentifier("https://contoso.vault.azure.net/secrets/aadSecret")
                    .withIdentityClientId("ceaa6b06-c00f-43ef-99ac-f53d1fe876a0"))
            .withSecret(true)
            .create();
    }
}
```
