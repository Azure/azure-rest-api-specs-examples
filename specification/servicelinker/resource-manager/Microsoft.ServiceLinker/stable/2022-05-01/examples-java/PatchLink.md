Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-servicelinker_1.0.0-beta.1/sdk/servicelinker/azure-resourcemanager-servicelinker/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.servicelinker.models.AzureResource;
import com.azure.resourcemanager.servicelinker.models.LinkerResource;
import com.azure.resourcemanager.servicelinker.models.ServicePrincipalSecretAuthInfo;

/** Samples for Linker Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/PatchLink.json
     */
    /**
     * Sample code: PatchLink.
     *
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void patchLink(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        LinkerResource resource =
            manager
                .linkers()
                .getWithResponse(
                    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app",
                    "linkName",
                    Context.NONE)
                .getValue();
        resource
            .update()
            .withTargetService(
                new AzureResource()
                    .withId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"))
            .withAuthInfo(
                new ServicePrincipalSecretAuthInfo().withClientId("name").withPrincipalId("id").withSecret("secret"))
            .apply();
    }
}
```
