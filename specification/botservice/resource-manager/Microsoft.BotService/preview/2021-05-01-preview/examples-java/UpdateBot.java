
import com.azure.resourcemanager.botservice.models.Bot;
import com.azure.resourcemanager.botservice.models.BotProperties;
import com.azure.resourcemanager.botservice.models.Kind;
import com.azure.resourcemanager.botservice.models.MsaAppType;
import com.azure.resourcemanager.botservice.models.PublicNetworkAccess;
import com.azure.resourcemanager.botservice.models.Sku;
import com.azure.resourcemanager.botservice.models.SkuName;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Bots Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/UpdateBot.json
     */
    /**
     * Sample code: Update Bot.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void updateBot(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        Bot resource = manager.bots()
            .getByResourceGroupWithResponse("OneResourceGroupName", "samplebotname", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).withProperties(new BotProperties()
            .withDisplayName("The Name of the bot").withDescription("The description of the bot")
            .withIconUrl("http://myicon").withEndpoint("http://mybot.coffee")
            .withMsaAppType(MsaAppType.USER_ASSIGNED_MSI).withMsaAppId("msaappid").withMsaAppTenantId("msaapptenantid")
            .withMsaAppMsiResourceId(
                "/subscriptions/foo/resourcegroups/bar/providers/microsoft.managedidentity/userassignedidentities/sampleId")
            .withDeveloperAppInsightKey("fakeTokenPlaceholder").withDeveloperAppInsightsApiKey("fakeTokenPlaceholder")
            .withDeveloperAppInsightsApplicationId("appinsightsappid")
            .withLuisAppIds(Arrays.asList("luisappid1", "luisappid2")).withLuisKey("fakeTokenPlaceholder")
            .withIsCmekEnabled(true).withCmekKeyVaultUrl("fakeTokenPlaceholder")
            .withPublicNetworkAccess(PublicNetworkAccess.ENABLED).withDisableLocalAuth(true)
            .withSchemaTransformationVersion("1.0")).withSku(new Sku().withName(SkuName.S1)).withKind(Kind.SDK)
            .withEtag("etag1").apply();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
