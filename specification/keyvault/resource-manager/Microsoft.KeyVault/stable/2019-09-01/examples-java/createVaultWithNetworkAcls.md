Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.keyvault.models.IpRule;
import com.azure.resourcemanager.keyvault.models.NetworkRuleAction;
import com.azure.resourcemanager.keyvault.models.NetworkRuleBypassOptions;
import com.azure.resourcemanager.keyvault.models.NetworkRuleSet;
import com.azure.resourcemanager.keyvault.models.Sku;
import com.azure.resourcemanager.keyvault.models.SkuFamily;
import com.azure.resourcemanager.keyvault.models.SkuName;
import com.azure.resourcemanager.keyvault.models.VaultCreateOrUpdateParameters;
import com.azure.resourcemanager.keyvault.models.VaultProperties;
import com.azure.resourcemanager.keyvault.models.VirtualNetworkRule;
import java.util.Arrays;
import java.util.UUID;

/** Samples for Vaults CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2019-09-01/examples/createVaultWithNetworkAcls.json
     */
    /**
     * Sample code: Create or update a vault with network acls.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAVaultWithNetworkAcls(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .vaults()
            .manager()
            .serviceClient()
            .getVaults()
            .createOrUpdate(
                "sample-resource-group",
                "sample-vault",
                new VaultCreateOrUpdateParameters()
                    .withLocation("westus")
                    .withProperties(
                        new VaultProperties()
                            .withTenantId(UUID.fromString("00000000-0000-0000-0000-000000000000"))
                            .withSku(new Sku().withFamily(SkuFamily.A).withName(SkuName.STANDARD))
                            .withEnabledForDeployment(true)
                            .withEnabledForDiskEncryption(true)
                            .withEnabledForTemplateDeployment(true)
                            .withNetworkAcls(
                                new NetworkRuleSet()
                                    .withBypass(NetworkRuleBypassOptions.AZURE_SERVICES)
                                    .withDefaultAction(NetworkRuleAction.DENY)
                                    .withIpRules(
                                        Arrays
                                            .asList(
                                                new IpRule().withValue("124.56.78.91"),
                                                new IpRule().withValue("'10.91.4.0/24'")))
                                    .withVirtualNetworkRules(
                                        Arrays
                                            .asList(
                                                new VirtualNetworkRule()
                                                    .withId(
                                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1"))))),
                Context.NONE);
    }
}
```
