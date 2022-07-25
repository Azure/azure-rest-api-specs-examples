import com.azure.resourcemanager.mobilenetwork.models.KeyVaultKey;
import com.azure.resourcemanager.mobilenetwork.models.MobileNetworkResourceId;

/** Samples for SimGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimGroupCreate.json
     */
    /**
     * Sample code: Create SIM group.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createSIMGroup(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .simGroups()
            .define("testSimGroup")
            .withRegion("eastus")
            .withExistingResourceGroup("rg1")
            .withEncryptionKey(new KeyVaultKey().withKeyUrl("https://contosovault.vault.azure.net/keys/azureKey"))
            .withMobileNetwork(
                new MobileNetworkResourceId()
                    .withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"))
            .create();
    }
}
