import com.azure.resourcemanager.desktopvirtualization.models.SendMessage;

/** Samples for UserSessions SendMessage. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/UserSession_SendMessage_Post.json
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
                com.azure.core.util.Context.NONE);
    }
}
