
import com.azure.resourcemanager.botservice.models.Key;
import com.azure.resourcemanager.botservice.models.RegenerateKeysChannelName;
import com.azure.resourcemanager.botservice.models.SiteInfo;

/**
 * Samples for DirectLine RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/
     * DirectlineRegenerateKeys.json
     */
    /**
     * Sample code: Regenerate Keys for DirectLine Channel Site.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void
        regenerateKeysForDirectLineChannelSite(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.directLines().regenerateKeysWithResponse("OneResourceGroupName", "samplebotname",
            RegenerateKeysChannelName.DIRECT_LINE_CHANNEL,
            new SiteInfo().withSiteName("testSiteName").withKey(Key.KEY1), com.azure.core.util.Context.NONE);
    }
}
