
import com.azure.resourcemanager.servicelinker.fluent.models.LinkerResourceInner;
import com.azure.resourcemanager.servicelinker.models.AzureResource;
import com.azure.resourcemanager.servicelinker.models.SecretAuthInfo;
import com.azure.resourcemanager.servicelinker.models.SecretStore;

/**
 * Samples for Connector CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * PutConnector.json
     */
    /**
     * Sample code: PutConnector.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void putConnector(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.connectors().createOrUpdate("00000000-0000-0000-0000-000000000000", "test-rg", "westus",
            "connectorName",
            new LinkerResourceInner().withTargetService(new AzureResource().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"))
                .withAuthInfo(new SecretAuthInfo()).withSecretStore(
                    new SecretStore().withKeyVaultId("fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
    }
}
