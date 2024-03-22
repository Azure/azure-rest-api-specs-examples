
import com.azure.resourcemanager.confluent.models.AccessInviteUserAccountModel;
import com.azure.resourcemanager.confluent.models.AccessInvitedUserDetails;

/**
 * Samples for Access InviteUser.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Access_InviteUser.json
     */
    /**
     * Sample code: Access_InviteUser.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void accessInviteUser(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.access().inviteUserWithResponse("myResourceGroup", "myOrganization",
            new AccessInviteUserAccountModel().withInvitedUserDetails(
                new AccessInvitedUserDetails().withInvitedEmail("user2@onmicrosoft.com").withAuthType("AUTH_TYPE_SSO")),
            com.azure.core.util.Context.NONE);
    }
}
