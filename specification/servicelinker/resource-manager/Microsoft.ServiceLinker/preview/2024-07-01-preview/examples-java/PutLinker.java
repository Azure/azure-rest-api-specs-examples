
import com.azure.resourcemanager.servicelinker.models.AzureResource;
import com.azure.resourcemanager.servicelinker.models.SecretAuthInfo;
import com.azure.resourcemanager.servicelinker.models.VNetSolution;
import com.azure.resourcemanager.servicelinker.models.VNetSolutionType;
import com.azure.resourcemanager.servicelinker.models.ValueSecretInfo;

/**
 * Samples for Linker CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * PutLinker.json
     */
    /**
     * Sample code: PutLinker.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void putLinker(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.linkers().define("linkName").withExistingResourceUri(
            "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app")
            .withTargetService(new AzureResource().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DBforPostgreSQL/servers/test-pg/databases/test-db"))
            .withAuthInfo(
                new SecretAuthInfo().withName("name").withSecretInfo(new ValueSecretInfo().withValue("secret")))
            .withVNetSolution(new VNetSolution().withType(VNetSolutionType.SERVICE_ENDPOINT)).create();
    }
}
