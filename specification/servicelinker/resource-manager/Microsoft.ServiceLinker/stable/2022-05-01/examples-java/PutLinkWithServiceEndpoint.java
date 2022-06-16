import com.azure.resourcemanager.servicelinker.models.AzureResource;
import com.azure.resourcemanager.servicelinker.models.KeyVaultSecretUriSecretInfo;
import com.azure.resourcemanager.servicelinker.models.SecretAuthInfo;
import com.azure.resourcemanager.servicelinker.models.VNetSolution;
import com.azure.resourcemanager.servicelinker.models.VNetSolutionType;

/** Samples for Linker CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/PutLinkWithServiceEndpoint.json
     */
    /**
     * Sample code: PutLinkWithServiceEndpoint.
     *
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void putLinkWithServiceEndpoint(
        com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager
            .linkers()
            .define("linkName")
            .withExistingResourceUri(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app")
            .withTargetService(
                new AzureResource()
                    .withId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DBforPostgreSQL/servers/test-pg/databases/test-db"))
            .withAuthInfo(
                new SecretAuthInfo()
                    .withName("name")
                    .withSecretInfo(
                        new KeyVaultSecretUriSecretInfo()
                            .withValue(
                                "https://vault-name.vault.azure.net/secrets/secret-name/00000000000000000000000000000000")))
            .withVNetSolution(new VNetSolution().withType(VNetSolutionType.SERVICE_ENDPOINT))
            .create();
    }
}
