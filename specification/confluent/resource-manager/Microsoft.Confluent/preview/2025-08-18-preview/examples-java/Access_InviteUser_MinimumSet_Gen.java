
import com.azure.resourcemanager.confluent.models.AccessInviteUserAccountModel;

/**
 * Samples for Access InviteUser.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Access_InviteUser_MinimumSet_Gen.json
     */
    /**
     * Sample code: Access_InviteUser_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void accessInviteUserMinimumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.access().inviteUserWithResponse("rgconfluent", "skqsedhorkejhhntdsiwroffkjld",
            new AccessInviteUserAccountModel(), com.azure.core.util.Context.NONE);
    }
}
