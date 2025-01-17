
import com.azure.resourcemanager.servicelinker.models.AzureResource;
import com.azure.resourcemanager.servicelinker.models.LinkerPatch;
import com.azure.resourcemanager.servicelinker.models.ServicePrincipalSecretAuthInfo;

/**
 * Samples for Connector Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * PatchConnector.json
     */
    /**
     * Sample code: PatchConnector.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void patchConnector(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.connectors().update("00000000-0000-0000-0000-000000000000", "test-rg", "westus", "connectorName",
            new LinkerPatch().withTargetService(new AzureResource().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"))
                .withAuthInfo(new ServicePrincipalSecretAuthInfo().withClientId("name").withPrincipalId("id")
                    .withSecret("fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
    }
}
