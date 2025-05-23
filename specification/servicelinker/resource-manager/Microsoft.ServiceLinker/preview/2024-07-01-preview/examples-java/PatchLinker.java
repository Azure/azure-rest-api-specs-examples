
import com.azure.resourcemanager.servicelinker.models.AzureResource;
import com.azure.resourcemanager.servicelinker.models.LinkerResource;
import com.azure.resourcemanager.servicelinker.models.ServicePrincipalSecretAuthInfo;

/**
 * Samples for Linker Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * PatchLinker.json
     */
    /**
     * Sample code: PatchLinker.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void patchLinker(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        LinkerResource resource = manager.linkers().getWithResponse(
            "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app",
            "linkName", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTargetService(new AzureResource().withId(
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"))
            .withAuthInfo(new ServicePrincipalSecretAuthInfo().withClientId("name").withPrincipalId("id")
                .withSecret("fakeTokenPlaceholder"))
            .apply();
    }
}
