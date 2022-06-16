import com.azure.core.util.Context;
import com.azure.resourcemanager.botservice.models.Key;
import com.azure.resourcemanager.botservice.models.RegenerateKeysChannelName;
import com.azure.resourcemanager.botservice.models.SiteInfo;

/** Samples for DirectLine RegenerateKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/WebChatRegenerateKeys.json
     */
    /**
     * Sample code: Regenerate Keys for WebChat Channel Site.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void regenerateKeysForWebChatChannelSite(
        com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .directLines()
            .regenerateKeysWithResponse(
                "OneResourceGroupName",
                "samplebotname",
                RegenerateKeysChannelName.WEB_CHAT_CHANNEL,
                new SiteInfo().withSiteName("testSiteName").withKey(Key.KEY1),
                Context.NONE);
    }
}
