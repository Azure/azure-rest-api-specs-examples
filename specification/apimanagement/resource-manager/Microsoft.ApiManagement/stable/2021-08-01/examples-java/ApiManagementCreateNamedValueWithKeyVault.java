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
