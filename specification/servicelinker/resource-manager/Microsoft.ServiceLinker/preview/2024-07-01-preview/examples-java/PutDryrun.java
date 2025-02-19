
import com.azure.resourcemanager.servicelinker.fluent.models.DryrunResourceInner;
import com.azure.resourcemanager.servicelinker.models.AzureResource;
import com.azure.resourcemanager.servicelinker.models.CreateOrUpdateDryrunParameters;
import com.azure.resourcemanager.servicelinker.models.SecretAuthInfo;
import com.azure.resourcemanager.servicelinker.models.ValueSecretInfo;

/**
 * Samples for LinkersOperation CreateDryrun.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * PutDryrun.json
     */
    /**
     * Sample code: PutDryrun.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void putDryrun(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.linkersOperations().createDryrun(
            "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app",
            "dryrunName",
            new DryrunResourceInner().withParameters(new CreateOrUpdateDryrunParameters()
                .withTargetService(new AzureResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"))
                .withAuthInfo(
                    new SecretAuthInfo().withName("name").withSecretInfo(new ValueSecretInfo().withValue("secret")))),
            com.azure.core.util.Context.NONE);
    }
}
