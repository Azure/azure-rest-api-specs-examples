
import com.azure.resourcemanager.servicelinker.models.AzureResource;
import com.azure.resourcemanager.servicelinker.models.CreateOrUpdateDryrunParameters;
import com.azure.resourcemanager.servicelinker.models.DryrunResource;
import com.azure.resourcemanager.servicelinker.models.SecretAuthInfo;
import com.azure.resourcemanager.servicelinker.models.ValueSecretInfo;

/**
 * Samples for Connector UpdateDryrun.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * ConnectorDryrunUpdate.json
     */
    /**
     * Sample code: ConnectorDryrunUpdate.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void connectorDryrunUpdate(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        DryrunResource resource = manager.connectors().getDryrunWithResponse("00000000-0000-0000-0000-000000000000",
            "test-rg", "westus", "dryrunName", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withParameters(new CreateOrUpdateDryrunParameters().withTargetService(new AzureResource().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"))
                .withAuthInfo(
                    new SecretAuthInfo().withName("name").withSecretInfo(new ValueSecretInfo().withValue("secret"))))
            .apply();
    }
}
