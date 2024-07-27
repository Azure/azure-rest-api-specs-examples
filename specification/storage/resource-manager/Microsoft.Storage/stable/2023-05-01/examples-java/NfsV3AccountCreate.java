
import com.azure.resourcemanager.storage.models.Bypass;
import com.azure.resourcemanager.storage.models.DefaultAction;
import com.azure.resourcemanager.storage.models.Kind;
import com.azure.resourcemanager.storage.models.NetworkRuleSet;
import com.azure.resourcemanager.storage.models.Sku;
import com.azure.resourcemanager.storage.models.SkuName;
import com.azure.resourcemanager.storage.models.StorageAccountCreateParameters;
import com.azure.resourcemanager.storage.models.VirtualNetworkRule;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageAccounts Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/NfsV3AccountCreate.json
     */
    /**
     * Sample code: NfsV3AccountCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nfsV3AccountCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts().create("res9101", "sto4445",
            new StorageAccountCreateParameters().withSku(new Sku().withName(SkuName.PREMIUM_LRS))
                .withKind(Kind.BLOCK_BLOB_STORAGE).withLocation("eastus")
                .withNetworkRuleSet(new NetworkRuleSet().withBypass(Bypass.AZURE_SERVICES)
                    .withVirtualNetworkRules(Arrays.asList(new VirtualNetworkRule().withVirtualNetworkResourceId(
                        "/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.Network/virtualNetworks/net123/subnets/subnet12")))
                    .withIpRules(Arrays.asList()).withDefaultAction(DefaultAction.ALLOW))
                .withEnableHttpsTrafficOnly(false).withEnableExtendedGroups(true).withIsHnsEnabled(true)
                .withEnableNfsV3(true),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
