
import com.azure.resourcemanager.confluent.models.AccessInviteUserAccountModel;
import com.azure.resourcemanager.confluent.models.AccessInvitedUserDetails;

/**
 * Samples for Access InviteUser.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Access_InviteUser_MaximumSet_Gen.json
     */
    /**
     * Sample code: Access_InviteUser_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void accessInviteUserMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.access().inviteUserWithResponse("rgconfluent", "aqwpihgldcvqwq",
            new AccessInviteUserAccountModel().withOrganizationId("aojvtivybqtuwwulokimwyh").withEmail("jtborwwroz")
                .withUpn("eyck").withInvitedUserDetails(new AccessInvitedUserDetails()
                    .withInvitedEmail("ozfkzouvjbvndqpyoxqbwtpzeiip").withAuthType("yaokrbtlql")),
            com.azure.core.util.Context.NONE);
    }
}
