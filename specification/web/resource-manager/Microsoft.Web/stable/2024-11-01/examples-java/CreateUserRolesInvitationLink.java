
import com.azure.resourcemanager.appservice.models.StaticSiteUserInvitationRequestResource;

/**
 * Samples for StaticSites CreateUserRolesInvitationLink.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/CreateUserRolesInvitationLink.json
     */
    /**
     * Sample code: Create an invitation link for a user for a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAnInvitationLinkForAUserForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().createUserRolesInvitationLinkWithResponse("rg",
            "testStaticSite0",
            new StaticSiteUserInvitationRequestResource().withDomain("happy-sea-15afae3e.azurestaticwebsites.net")
                .withProvider("aad").withUserDetails("username").withRoles("admin,contributor")
                .withNumHoursToExpiration(1),
            com.azure.core.util.Context.NONE);
    }
}
