import com.azure.core.util.Context;
import com.azure.resourcemanager.labservices.models.InviteBody;

/** Samples for Users Invite. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Users/inviteUser.json
     */
    /**
     * Sample code: inviteUser.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void inviteUser(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager
            .users()
            .invite(
                "testrg123",
                "testlab",
                "testuser",
                new InviteBody().withText("Invitation to lab testlab"),
                Context.NONE);
    }
}
