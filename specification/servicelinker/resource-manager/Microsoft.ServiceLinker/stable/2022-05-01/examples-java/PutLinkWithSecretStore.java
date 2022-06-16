import com.azure.resourcemanager.servicelinker.models.AzureResource;
import com.azure.resourcemanager.servicelinker.models.SecretAuthInfo;
import com.azure.resourcemanager.servicelinker.models.SecretStore;

/** Samples for Linker CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/PutLinkWithSecretStore.json
     */
    /**
     * Sample code: PutLinkWithSecretStore.
     *
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void putLinkWithSecretStore(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager
            .linkers()
            .define("linkName")
            .withExistingResourceUri(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app")
            .withTargetService(
                new AzureResource()
                    .withId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"))
            .withAuthInfo(new SecretAuthInfo())
            .withSecretStore(
                new SecretStore()
                    .withKeyVaultId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.KeyVault/vaults/test-kv"))
            .create();
    }
}
