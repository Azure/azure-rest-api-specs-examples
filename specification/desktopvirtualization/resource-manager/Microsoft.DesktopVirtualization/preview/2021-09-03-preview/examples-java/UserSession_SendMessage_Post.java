import com.azure.core.util.Context;
import com.azure.resourcemanager.desktopvirtualization.models.SendMessage;

/** Samples for UserSessions SendMessage. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/UserSession_SendMessage_Post.json
     */
    /**
     * Sample code: UserSession_SendMessage_Post.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void userSessionSendMessagePost(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .userSessions()
            .sendMessageWithResponse(
                "resourceGroup1",
                "hostPool1",
                "sessionHost1.microsoft.com",
                "1",
                new SendMessage().withMessageTitle("title").withMessageBody("body"),
                Context.NONE);
    }
}
