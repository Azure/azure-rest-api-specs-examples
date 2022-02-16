Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.models.StaticSiteUserInvitationRequestResource;

/** Samples for StaticSites CreateUserRolesInvitationLink. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateUserRolesInvitationLink.json
     */
    /**
     * Sample code: Create an invitation link for a user for a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAnInvitationLinkForAUserForAStaticSite(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .createUserRolesInvitationLinkWithResponse(
                "rg",
                "testStaticSite0",
                new StaticSiteUserInvitationRequestResource()
                    .withDomain("happy-sea-15afae3e.azurestaticwebsites.net")
                    .withProvider("aad")
                    .withUserDetails("username")
                    .withRoles("admin,contributor")
                    .withNumHoursToExpiration(1),
                Context.NONE);
    }
}
```
