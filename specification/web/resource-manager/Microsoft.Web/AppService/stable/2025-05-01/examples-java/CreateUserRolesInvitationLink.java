
import com.azure.resourcemanager.appservice.models.StaticSiteUserInvitationRequestResource;

/**
 * Samples for StaticSites CreateUserRolesInvitationLink.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/CreateUserRolesInvitationLink.json
     */
    /**
     * Sample code: Create an invitation link for a user for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        createAnInvitationLinkForAUserForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().createUserRolesInvitationLinkWithResponse("rg", "testStaticSite0",
            new StaticSiteUserInvitationRequestResource().withDomain("happy-sea-15afae3e.azurestaticwebsites.net")
                .withProvider("aad").withUserDetails("username").withRoles("admin,contributor")
                .withNumHoursToExpiration(1),
            com.azure.core.util.Context.NONE);
    }
}
